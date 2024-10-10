ARG RHEL_VERSION=9

FROM quay.io/centos/centos:stream$RHEL_VERSION as IMAGE

ARG RHEL_VERSION
ARG release=main
ARG version=latest

LABEL com.redhat.component="assisted-service"
LABEL description="Service that installs Openshift"
LABEL summary="Service that installs Openshift"
LABEL io.k8s.description="Service that installs Openshift"
LABEL distribution-scope="public"
LABEL name="assisted-service"
LABEL release="${release}"
LABEL version="${version}"
LABEL url="https://github.com/openshift/assisted-service"
LABEL vendor="Red Hat, Inc."
LABEL maintainer="Red Hat"

# Ensure UID can write in data dir (e.g.: when using podman, docker, ...)
# Ensure root group can write in data dir when deployed on OCP
# https://docs.openshift.com/container-platform/4.16/openshift_images/create-images.html#use-uid_create-images
ARG UID=1001
ARG GID=0
RUN mkdir /data && chmod 775 /data && chown $UID:$GID /data

ENV GODEBUG=madvdontneed=1
ENV GOGC=50

COPY hack/container_build_scripts/replace_dnf_repositories_ref_if_needed.sh /usr/local/bin

# multiarch fix WRKLDS-222 and https://bugzilla.redhat.com/show_bug.cgi?id=2111537
# openshift-install requires nmstate nmstate-libs
RUN replace_dnf_repositories_ref_if_needed.sh $RHEL_VERSION && \
    dnf install -y --setopt=install_weak_deps=False --setopt=fastestmirror=1 \
    skopeo libvirt-libs nmstate nmstate-libs libksba libxml2 && \
    dnf clean all

FROM registry.access.redhat.com/ubi$RHEL_VERSION/go-toolset:1.21 AS GOLANG

WORKDIR /app
USER root

ENV GOROOT=/usr/lib/golang
ENV GOPATH=/opt/app-root/src/go
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

COPY --from=IMAGE /usr/local/bin/replace_dnf_repositories_ref_if_needed.sh  /usr/local/bin/
RUN go install github.com/google/go-licenses@v1.6.0

# installing build TOOLS
RUN replace_dnf_repositories_ref_if_needed.sh $RHEL_VERSION  && \
    dnf install --setopt=install_weak_deps=False --setopt=fastestmirror=1 -y nmstate-devel  && \
    dnf clean all

FROM GOLANG as BUILDER

ENV CGO_ENABLED=1
ENV GOFLAGS=""
ENV GO111MODULE=on

COPY . .
RUN go mod download

RUN go-licenses save --save_path /build/licenses/ ./...

WORKDIR /app/cmd

RUN go build -o /build/assisted-service

RUN cd operator && \
    go build -o /build/assisted-service-operator
RUN cd webadmission &&  \
    go build -o /build/assisted-service-admission
RUN cd agentbasedinstaller/client && \
    go build -o /build/agent-installer-client

FROM IMAGE as FINAL

COPY --from=BUILDER /build/licenses /licenses
COPY --from=BUILDER /build/assisted-service /assisted-service
COPY --from=BUILDER /build/assisted-service-operator /assisted-service-operator
COPY --from=BUILDER /build/assisted-service-admission /assisted-service-admission
COPY --from=BUILDER /build/agent-installer-client /usr/local/bin/agent-installer-client
RUN ln -s /usr/local/bin/agent-installer-client /agent-based-installer-register-cluster-and-infraenv

USER $UID:$GID

CMD ["/assisted-service"]