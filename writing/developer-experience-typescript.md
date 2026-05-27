---
title: 'TypeScript has a terrible developer experience'
date: '2024-2-21'
tags: ['Typescript', 'developer experience', 'node']
summary: 'TypeScript has a terrible developer experience tools when compared to go'
---

I find the development process with TypeScript painful in every step of development due to divergent tooling. I find the TypeScript language good, but I find everything around the language bad. In comparsion, I find the Go language from easy in basically every aspect (Check my other most...).

You can experience this headache setting up a basic webserver, as you go through a lot of the steps in the development process. So here we go...

How fast can you deploy a server in TypeScript from scratch?

3... 2... 1... go!

### 1. Make a folder:

```
mkdir typescript-deployment-minimal
```

### 2. initiate npm

```shell
npm init -y
```

This creates a lot of values I don't care for like version, description, main, author and license.

### 3. initiate typescript

```shell
npm install typescript @types/node --save-dev
tsc --init # create tsconfig.json
```

This creates the tsconfig.json, required for Typescript.

### 4. get stuck looking at tsconfig.json

The `tsconfig.json` file generated from `tsc --init` is 109 lines long, and has a doesn't have sane defaults...

This is a sane default to have though (I had to look this up on youtube, the documentation isn't great):

```json
{
  "compilerOptions": {
    "target": "esNext",
    "module": "NodeNext",
    "rootDir": "./src",
    "allowJs": false,
    "sourceMap": true,
    "outDir": "./dist",
    "esModuleInterop": true,
    "forceConsistentCasingInFileNames": true,
    "strict": true,
    "skipLibCheck": true
  }
}
```

It's got a `src` for code and `dist` folder for compiled code (or javascript). You probably don't need `"sourceMap": true,`

### 5. write a web server in src

Let's use fastify as the web server as people would laugh at you if you just use the http module, and express.js is legacy now (no http/2).

```shell
npm install fastify
```

And let's hit chatgpt to quickly get some code:

```Typescript
import fastify from 'fastify';

const server = fastify({
    logger: true
});

server.get('/', async (request, reply) => {
    return { hello: 'world' };
});

const start = async () => {
    try {
        await server.listen({ port: 3000 });
    } catch (err) {
        server.log.error(err);
        process.exit(1);
    }
};

start();

```

This is great. It's short and we have a logger too. Otherwise we would have had to have installed and configured a logger like `winston`.

I wanted to protect you from this, but here is the type of the fastify server:

```Typescript
declare function fastify<
  Server extends http.Server,
  Request extends RawRequestDefaultExpression<Server> = RawRequestDefaultExpression<Server>,
  Reply extends RawReplyDefaultExpression<Server> = RawReplyDefaultExpression<Server>,
  Logger extends FastifyBaseLogger = FastifyBaseLogger,
  TypeProvider extends FastifyTypeProvider = FastifyTypeProviderDefault,
>(opts?: fastify.FastifyHttpOptions<Server, Logger>): FastifyInstance<Server, Request, Reply, Logger, TypeProvider> & PromiseLike<FastifyInstance<Server, Request, Reply, Logger, TypeProvider>>
```

To get the port of the server it's:

```Typescript
        const address = server.server.address() as AddressInfo
        address.port
```

I have a lost a lot of time to complicated types like these.

### compile and run

compile the code with:

```shell
tsc
```

With bigger projects you probably want to move to swc, or the `speedy web compile` for builds as compiling is a lot faster. There is esbuild aswell.

run the code with:

```shell
node dist/index.js
```

Sweet...

With more advanced projects you probably want to compile with `swc` the speedy web compiler, or esbuild. I think they are a lot faster they `tsc` or typescript compile.

As an aside to run Typescript code with bun, it is just:

```shell
bun run index.ts
```

You are generally stepping through two tools to run and compile.

### deploy it - it's not done until it is running in production

#### setup the server

Let's create a tiny EC2 instance on amazon web services (AWS), like t4g.nano spot instance. We will just use ssh, and scp to setup the server.

Let's connect to the server

```shell
ssh -i ~/.ssh/ec2-key.pem ec2-user@52.64.95.107
```

Once on the server we need to install node. You would think it is just:

```shell
sudo yum install node
```

But that doesn't work.

It's actually:

```shell
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash # install node version manager
source ~/.bashrc # load nvm
nvm install --lts # install long term support version
```

Let's make a directory for the code as well on the server:

```shell
mkdir node
```

#### move the code to the server, and run the code

Now the server is setup let's move the code there.

```shell
# copy the package.json file for dependencies
scp -i ~/.ssh/ec2-key.pem package.json ec2-user@52.64.95.107:/home/ec2-user/node
# copy the package-lock.json for dependency versions
scp -i ~/.ssh/ec2-key.pem package-lock.json ec2-user@52.64.95.107:/home/ec2-user/node
# copy the dist directory for the compiled code
scp -i ~/.ssh/ec2-key.pem -r dist/ ec2-user@52.64.95.107:/home/ec2-user/node
```

Back on the server we need to re-install the packages, and run the code:

```shell
cd node
npm install --production # you could copy the node_modules, but i think installing is more safer
# the production flag seem you don't install devDependencies
# now let's run the server:
nohup node dist/index.js > application.log &
```

`nohup` mean no hang up, it prevents the linux server from shutting now the process. `&` runs the process in the background.

### thoughts on developer experience

And that's it, we have a server listening on port 3000, serving content.

I can deploy a web server somewhat quickly, but I felt pain on basically every step:

- The package.json file has a lot of garbage in it, which annoys me
- I get overwhelmed by the tsconfig.json file being 100 lines of options from `tsc --init`
- The types of some objects are crazy.
- Compiling can run into issues with misconfigured tsconfig.json, and with bigger projects you would want to move to the speedy web compiler `swc` for a lot more speed
- on the server I needed to install node, a node version manager and the libraries again, and scp a directory
- writing a test you need another tool (see below)

We also missed out on setting up prettier, and eslint, and sharing code between projects, and workspaces, which can be very confusing.

I did enjoy the good library support though.

Check my other post on how this compares to Go language.

### bonus... setting up testing

Here is a bonus of setting up tests with node. You can use jest, or vitest. I'm using vitest as it's more modern.

```shell
npm install -D vitest
```

You generally need `vitest.config.ts` file, and to add the test script to package.json.

From there you write tests in the file like `.test.ts`.

### References:

1. How to install node on amazon linux:
   https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/setting-up-node-on-ec2-instance.html
