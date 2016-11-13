FROM blang/golang-alpine
ADD api /
CMD ["/api"]
EXPOSE 8000
