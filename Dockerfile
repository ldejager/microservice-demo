FROM golang:1.7
ADD binary/api-linux-x86_64 /
CMD ["/api-linux-x86_64"]
EXPOSE 8000
