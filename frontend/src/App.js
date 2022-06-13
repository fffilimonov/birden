import {useState} from 'react'

import './App.css';

import axios from 'axios'

import Web3 from "web3";
import {ethers} from 'ethers'

import Contract from './Birden.json'

import ERC20 from './ERC20.json'

import Home from './pages/Home'

function App() {
  const [w3, setW3] = useState(undefined)
  const [address, setAddress] = useState('')
  const [tokenID, setTokenID] = useState(0)

  async function connect() {
      console.log(window.ethereum)
        const accs = await window.ethereum.request({method: 'eth_requestAccounts'})
        const EthersProvider = new ethers.providers.Web3Provider(window.ethereum);
        const web3 = new Web3(EthersProvider.provider);
        setW3(web3)
        setAddress(accs[0])
  }


  async function getToken() {

      const req = {
        wallet: address,
        tokenID: parseInt(tokenID),
      }

      const response = await axios.post(`/token`, req)
      let token = response.data.token.slice(1, -1);
      console.log("https://birden.mypinata.cloud/ipfs/bafybeidvt4trjo7kdi6ofyiah6626lzssdc2244zit6lh5vsg4ndtt5hea?accessToken="+token)
  }

  async function rent() {

      const celo = new w3.eth.Contract(
        ERC20.abi,
        '0xf194afdf50b03e69bd7d057c1aa9e10c9954e4c9',
      )
      const tx0 = await celo.methods.approve(Contract.networks['44787'].address,Web3.utils.toWei("1")).send({gasLimit: 10000000, from: address})
      console.log(tx0)

      const birden = new w3.eth.Contract(
        Contract.abi,
        Contract.networks['44787'].address,
      )
      const tx1 = await birden.methods.rent(tokenID).send({gasLimit: 10000000, from: address})
      console.log(tx1)
  }

  return (
      <div>
        <button onClick={connect}>
          Connect
        </button>
        <button onClick={getToken}>
          Token
        </button>
        <button onClick={rent}>
          Rent
        </button>
        <input
          type='number'
          placeholder='number'
          onChange={(e) => setTokenID(e.target.value)}
        />
          <Home />
      </div>
  );
}

export default App;
