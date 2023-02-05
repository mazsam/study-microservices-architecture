
const PROTO_PATH = __dirname + '/../protos/auth.proto';

const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });

const auth_proto = grpc.loadPackageDefinition(packageDefinition).auth;


const bcrypt = require("bcrypt")
const jwt = require('jsonwebtoken')

const config = require("./config")
const db = require('monk')(config.db.uri)
db.addMiddleware(require('monk-middleware-debug'))

/**
 * Implements the SayHello RPC method.
 */
async function register(call, res) {

    const users = db.get('users')

    try {
        const hash = await bcrypt.hash(call.request.password, parseInt(config.salt))
        const result = await users.insert({
            username: call.request.username,
            password: hash
        })
        res(null, { message: `Register successfully with id ${result._id}`, code: 201 });
    } catch (error) {
        res({ message: error, code: 400 });
    }

}

async function login(req, res) {
    const users = db.get('users')

    try {
        const user = await users.findOne({ username: req.request.username })

        if (!user) {
            res({
                message: "user not found",
                code: grpc.status.FAILED_PRECONDITION,
            })
        }
        const matchPassword = await bcrypt.compare(req.request.password, user.password)
        if (!matchPassword) {
            res({
                message: "password not match    ",
                code: grpc.status.FAILED_PRECONDITION,
            })
        }
        if (matchPassword) {
            const token = jwt.sign({
                sub: user._id,
                name: user.username,
            }, config.jwt.secret)


            res(null, {
                message: "success",
                data: {
                    username: user.username,
                    user_id: user._id,
                    token: token,
                },
                code: grpc.status.OK,
            })
        }


    } catch (error) {
        res({
            message: error,
            code: grpc.status.FAILED_PRECONDITION,
        })
    }
}

function introspection(req, res) {
    try {
        const verify = jwt.verify(req.request.token, config.jwt.secret)

        if (verify) {
            res(null, {
                message: "success",
                code: grpc.status.OK,
                data: {
                    username: verify.name,
                    user_id: verify.sub,
                    token: req.request.token
                }
            })
        }
    } catch (error) {
        res({
            message: "UNAUTHENTICATED",
            code: grpc.status.UNAUTHENTICATED
        })
    }

}
/**
 * Starts an RPC server that receives requests for the Greeter service at the
 * sample server port
 */
function main() {
    const server = new grpc.Server();
    server.addService(auth_proto.Auth.service, {
        register,
        login,
        introspection
    });
    server.bindAsync('0.0.0.0:50051', grpc.ServerCredentials.createInsecure(), () => {
        server.start();
    });


}

main();