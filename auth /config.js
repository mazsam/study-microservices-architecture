require('dotenv').config()

module.exports = {
    db: {
        uri: process.env.DB_URI
    },
    jwt: {
        secret: process.env.JWT_SECRET
    },
    salt: process.env.SALT
}