FROM gcr.io/distroless/base
COPY bins /
ENTRYPOINT ["/bins"] 