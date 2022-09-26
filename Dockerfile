FROM debian 
RUN apt update
RUN apt install golang wget unzip cmake make
RUN wget https://github.com/opencv/opencv/archive/4.5.5.zip
RUN unzip ./4.5.5.zip 
RUN cd opencv-4.5.5/ 
RUN cmake -S . -B build & make & make install
COPY libttvfaceengine1.so /lib/
COPY ttv_chech /usr/bin/ttv_chech
COPY  dict   /opt/ttv_dict

CMD [ "ttv_chech" ]

