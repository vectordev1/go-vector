# adjusted from https://github.com/vector/cpp-vector/blob/develop/docker/Dockerfile
FROM ubuntu:14.04

ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update
RUN apt-get upgrade -y

# Vector dependencies
RUN apt-get install -qy build-essential g++-4.8 git cmake libboost-all-dev libcurl4-openssl-dev wget
RUN apt-get install -qy automake unzip libgmp-dev libtool libleveldb-dev yasm libminiupnpc-dev libreadline-dev scons
RUN apt-get install -qy libjsoncpp-dev libargtable2-dev

# NCurses based GUI (not optional though for a succesful compilation, see https://github.com/vector/cpp-vector/issues/452 )
RUN apt-get install -qy libncurses5-dev

# Qt-based GUI
# RUN apt-get install -qy qtbase5-dev qt5-default qtdeclarative5-dev libqt5webkit5-dev

# Vector PPA
RUN apt-get install -qy software-properties-common
RUN add-apt-repository ppa:vector/vector
RUN apt-get update
RUN apt-get install -qy libcryptopp-dev libjson-rpc-cpp-dev

# Build Vector (HEADLESS)
RUN git clone --depth=1 --branch develop https://github.com/vector/cpp-vector
RUN mkdir -p cpp-vector/build
RUN cd cpp-vector/build && cmake .. -DCMAKE_BUILD_TYPE=Release -DHEADLESS=1 && make -j $(cat /proc/cpuinfo | grep processor | wc -l) && make install
RUN ldconfig

ENTRYPOINT ["/cpp-vector/build/test/createRandomTest"]

