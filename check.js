const ContractKit = require('@celo/contractkit');
const Web3 = require('web3');
const Birden = require('./abi/Birden.json');

require('dotenv').config();

const main = async () => {
  const web3 = new Web3('https://alfajores-forno.celo-testnet.org');
  const client = ContractKit.newKitFromWeb3(web3);

  const account = web3.eth.accounts.privateKeyToAccount(process.env.PRIVATE_KEY);
  const networkId = await web3.eth.net.getId();
  client.connection.addAccount(account.privateKey);
  const deployedNetwork = Birden.networks[networkId];

  let goldtoken = await client.contracts.getGoldToken()
  let celoBalance = await goldtoken.balanceOf(account.address)

    console.log(`Your account address: ${account.address}`)
    console.log(`Your account CELO balance: ${celoBalance.toString()}`)

  if (!deployedNetwork) {
    throw new Error(`${networkId} is not valid`);
  }

  let instance = new web3.eth.Contract(
    Birden.abi,
    deployedNetwork.address
    );

  const txObject = await instance.methods.safeMint('https://birden.mypinata.cloud/ipfs/bafybeidvt4trjo7kdi6ofyiah6626lzssdc2244zit6lh5vsg4ndtt5hea', goldtoken.address);
  let tx = await client.sendTransactionObject(txObject, { from: account.address })

  let receipt = await tx.waitReceipt()
   console.log(receipt)

  let tokens = await instance.methods.tokensInfo().call();
  console.log(tokens)

  let nftBalance = await instance.methods.balanceOf(account.address).call();
  console.log(nftBalance)
};

main().catch((err) => {
  console.error(err);
});

