const openpgp = require('openpgp');

const TestPrivateKey =
    `-----BEGIN PGP PRIVATE KEY BLOCK-----

-----END PGP PRIVATE KEY BLOCK-----`

const TestPublicKey =
    `-----BEGIN PGP PUBLIC KEY BLOCK-----

-----END PGP PUBLIC KEY BLOCK-----`

async function asyncCall() {
    try {
        const publicKeys = await openpgp.key.readArmored(TestPublicKey)
        const encrypted = await openpgp.encrypt({
            message: openpgp.message.fromText("hello"),
            publicKeys: publicKeys.keys,
        });
        console.log(encrypted)

        const message = await openpgp.message.readArmored(encrypted.data)

        const privateKey = (await openpgp.key.readArmored([TestPrivateKey]))
        const decrypt = await openpgp.decrypt({
            message: message,
            privateKeys: privateKey.keys,
        });
        console.log(decrypt.data)
    } catch (e) {
        console.error(e);
    }
}

asyncCall()



