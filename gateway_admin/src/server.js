const {ApolloServer} = require('apollo-server');
const {ApolloServerPluginLandingPageGraphQLPlayground} = require('apollo-server-core');
const {ApolloGateway, RemoteGraphQLDataSource} = require("@apollo/gateway");
const {ApolloLogPlugin} = require('apollo-log');
const {ApolloArmor} = require('@escape.tech/graphql-armor');

const axios = require('axios');
const {GraphQLError} = require("graphql/error");
const crypto = require('crypto');

console.log('service url', process.env.AUTH_ADMIN_URL)

class AuthenticatedDataSource extends RemoteGraphQLDataSource {

    async willSendRequest({request, context}) {
        
        request.http.headers.set('Cookie', context.cookies)
		
    }

    async didReceiveResponse({response, request, context}) {
        return response;
    }
}

const blockFieldSuggestion = (process.env.BLOCK_FIELD_SUGGESTION === 'true');
const enableStackTrace = (process.env.ENABLE_STACK_TRACE === 'true');
const enableIntrospection = (process.env.ENABLE_INTROSPECTION === 'true');
const maxTokens = process.env.MAX_TOKENS;

const armor = new ApolloArmor({
    blockFieldSuggestion: {
        enabled: blockFieldSuggestion
    },
    maxDepth: {
        n: 20
    },
    costLimit: {
        maxCost: 300000
    },
    maxTokens: {
        n: maxTokens
    }
});

const protection = armor.protect()

const gateway = new ApolloGateway({
  experimental_pollInterval: process.env.POOL_TIME || 30000,
  serviceList: [
        {name: 'car_admin', url: `${ process.env.CAR_ADMIN_URL}/query`},
        {name: 'dealership_admin', url: `${ process.env.DEALERSHIP_ADMIN_URL}/query`},
  ],

  buildService({name, url}) {
    return new AuthenticatedDataSource({url});
  },
});
// Pass the ApolloGateway to the ApolloServer constructor
console.log("CORS SETTINGS DEV?:", process.env.DEV, "CORS_WHITELIST", process.env.CORS_WHITELIST)
const server = new ApolloServer({
  ...protection,
  debug: enableStackTrace,
  introspection: enableIntrospection,
  gateway,
  engine: {
    apiKey: process.env.APOLLO_METRICS_KEY,
    // reportSchema: true,
    // graphVariant: "current",
  },
  cors: {
    "origin": process.env.CORS_WHITELIST.replace(/\s/g, "").split(","),
    "credentials": true,
    "methods": "GET,HEAD,PUT,PATCH,POST,DELETE,OPTIONS",
    "preflightContinue": false,
    "optionsSuccessStatus": 204,
  },
    context: ({req, res}) => {
        return {
            cookies: req.headers.cookie,
            authorization: req.headers.authorization,
            origin: req.headers.origin,
            expressRes: res,
            refresh: false
        }
    },
  validationRules: protection.validationRules,
  plugins: [
        ...protection.plugins,
      {
        requestDidStart(requestContext) {
          let startDate = Date.now();
         // console.log(startDate, "requestContext", requestContext.request.http.headers.get("user-agent"))
         // console.log("authorization", requestContext.request.http.headers.get("Authorization"))
          // console.log('Request started! Query:\n' +
          //   requestContext.request.query);
          return {
            willSendResponse({context, response}) {
              // Append our final result to the outgoing response headers
              let diff = Date.now() - startDate;
              // console.log(startDate, diff, "RESPONSE");
              response.http.headers.set(
                'Server-Timing',
                startDate
              );
            }
          };
        }
      },
    // ApolloLogPlugin({timestamp: true})
  ApolloServerPluginLandingPageGraphQLPlayground({
    settings: {
      "request.credentials": "include"
    }
  })

  ],
  subscriptions: false,
});


server.listen({port: process.env.GATEWAY_PORT || 3000}).then(({url}) => {
  console.log(`🚀 Server ready at ${url}`);
});
