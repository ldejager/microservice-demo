FROM alpine
ADD api /
CMD ["/api"]
EXPOSE 8000
