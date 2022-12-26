FROM node:19-alpine
RUN apk update && apk add bash
RUN apk add openjdk11

RUN npm install -g firebase-tools

WORKDIR /firestore
COPY firebase.json .

ENV FIREBASE_PROJECT_ID=desafio-c0479
ENV ENABLE_UI=true

EXPOSE 9090
EXPOSE 4000
EXPOSE 4400


CMD ["sh", "-c", "firebase --project=$FIREBASE_PROJECT_ID emulators:start --only firestore" ]