#   __                          __
#  /\ \                        /\ \
#  \_\ \  __  __    ___     ___\ \ \___      __      ___
#  /'_` \/\ \/\ \ /' _ `\  /'___\ \  _ `\  /'__`\  /' _ `\
# /\ \L\ \ \ \_\ \/\ \/\ \/\ \__/\ \ \ \ \/\ \L\.\_/\ \/\ \
# \ \___,_\/`____ \ \_\ \_\ \____\\ \_\ \_\ \__/.\_\ \_\ \_\
#  \/__,_ /`/___/> \/_/\/_/\/____/ \/_/\/_/\/__/\/_/\/_/\/_/
#             /\___/
#             \/__/

FROM golang:latest
MAINTAINER John Lettman "the@johnlettman.com"

# Expose the Dynchan default port.
EXPOSE 8080

# Move the Dynchan sources onto the container.
RUN mkdir -p /go/src/github.com/jlettman/dynchan
ADD . /go/src/github.com/jlettman/dynchan

# Get the Dynchan dependencies.
RUN go get github.com/jlettman/dynchan/...

# Install Dynchan.
RUN go install github.com/jlettman/dynchan

# Entry.
CMD ["dynchan"]
