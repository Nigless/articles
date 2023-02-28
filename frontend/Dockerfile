FROM node

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm i

CMD ["/bin/sh" , "-c", "npx vite"]