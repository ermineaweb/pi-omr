FROM gradle:jdk17

WORKDIR /
RUN git clone https://github.com/Audiveris/audiveris.git

WORKDIR /audiveris
RUN git checkout 5.2.5
RUN apt-get update && apt-get install -y tesseract-ocr tesseract-ocr-ita
RUN gradle wrapper --stacktrace
RUN ./gradlew clean build
RUN tar -xf build/distributions/Audiveris-5.3-alpha.tar

# other requirements
RUN apt-get update && apt-get install -y lilypond libsndfile-dev

WORKDIR /data

CMD "/audiveris/Audiveris-5.3-alpha/bin/Audiveris"
