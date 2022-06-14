import styles from './Home.module.css'
import classNames from 'classnames';
import img from '../assets/images/img.png'
import {useEffect, useRef, useState} from 'react';
import {WithdrawChartView} from '../components/WithdrawChartVew';
import {ReactComponent as MenuIcon} from '../assets/images/menu.svg';
import {ReactComponent as CoinIcon} from '../assets/images/coin.svg';
import {ethers} from 'ethers';
import Web3 from 'web3';
import axios from 'axios';
import ERC20 from '../ERC20.json';
import Contract from '../Birden.json';

const Home = () => {
    const tokenID = window.location.search ? window.location.search.split('?')[1] : 4
    const pricePerSecond = 1/600
    const timerLength = 600

    const [open, setOpen] = useState(true)
    const [videoUrl, setVideoUrl] = useState();
    const [playing, setPlaying] = useState(false)
    const [w3, setW3] = useState(undefined)
    const [address, setAddress] = useState('')
    const [balance, setBalance] = useState(0)
    const [timeLeft, setTimeLeft] = useState(timerLength)

    const balanceLeft = useRef(0)
    const timeLeftRef = useRef(timerLength)
    const timer = useRef()
    const videoRef = useRef()

    useEffect(() => {
        void connect();

        return () => {
            if (timer.current) {
                clearInterval(timer.current)
            }
        }
    }, [])

    async function connect() {
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

        const response = await axios.post(`/token`, req).catch(() => {
            rent()
        })
        let token = response.data.token.slice(1, -1);
        setVideoUrl('https://birden.mypinata.cloud/ipfs/bafybeidvt4trjo7kdi6ofyiah6626lzssdc2244zit6lh5vsg4ndtt5hea?accessToken=' + token)
        setBalance(1)
    }

    async function rent() {
        const celo = new w3.eth.Contract(
            ERC20.abi,
            '0xf194afdf50b03e69bd7d057c1aa9e10c9954e4c9',
        )
        const tx0 = await celo.methods.approve(Contract.networks['44787'].address, Web3.utils.toWei('1')).send({
            gasLimit: 10000000,
            from: address
        })
        console.log(tx0)

        const birden = new w3.eth.Contract(
            Contract.abi,
            Contract.networks['44787'].address,
        )
        const tx1 = await birden.methods.rent(tokenID).send({gasLimit: 10000000, from: address})
        console.log(tx1)

        await getToken()
    }

    useEffect(() => {
        balanceLeft.current = balance - pricePerSecond * (timerLength - timeLeft)
        timeLeftRef.current = timeLeft
    }, [timeLeft, balance, pricePerSecond])

    const onPlayClick = (event) => {
        event.stopPropagation();

        if (videoUrl && videoRef.current?.paused) {
            videoRef.current.play();
            setPlaying(true);
            timer.current = setInterval(() => {
                if (timeLeftRef.current <= 0 || balanceLeft.current < 0) {
                    clearInterval(timer.current)
                    setPlaying(false)
                    setVideoUrl(null)
                    setTimeLeft(timerLength)
                    if (videoRef.current) {
                        videoRef.current.pause()
                    }
                } else {
                    setTimeLeft(prev => prev - 1)
                }
            }, 1000)
        }
    }

    return <div className={styles.container}>
        <div className={classNames(styles.leftSide, open && styles.active)}>
            <div className={styles.leftSideContent}>
                <span className={styles.leftSideTitle}>Chris Luno</span>
                <span className={styles.leftSideText}>
                    I'm glad you found me<br/>
                    Here, have some music <br/>
                    And drink more water <br/>
                    <br/>
                    ↓ ↓ ↓
                    <br/>
                    <strong>YouTube</strong>
                </span>


                <div className={styles.leftSideBadge}>
                    <div>
                        <span className={styles.leftSideBadgeTitle}>good morning house mix</span>
                        <span className={styles.leftSideBadgeText}>birden.io/ChrisLuno/goodmorning...</span>
                    </div>
                    <div className={styles.leftSideCoinBadge}><CoinIcon style={{marginRight: '4px'}}/> 1.00</div>
                </div>
            </div>
        </div>

        <div className={classNames(styles.rightSide, !open && styles.rightSideActive)}>
            <div className={styles.rightSideHeader}>
                <MenuIcon className={styles.righSideMenuIcon} onClick={() => setOpen(prev => !prev)}/>
                {address && <div className={styles.rightSideWallet}>{address.substring(address.length - 4)}</div>}
            </div>
            <div style={{position: 'relative', width: '100%'}}>
                {!videoUrl ?
                    <img className={styles.rightSideMedia} src={img} alt={'preview'}/> :
                    <video preload={'metadata'} controls={false} ref={videoRef} className={styles.rightSideMedia}
                           src={videoUrl} playsInline/>
                }

                {!playing && (
                    <svg onClick={onPlayClick}
                         style={{position: 'absolute', margin: 'auto', inset: 0, cursor: 'pointer'}} width="99"
                         height="99"
                         viewBox="0 0 99 99"
                         fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path
                            d="M49.374 98.2598C76.0752 98.2598 98.1826 76.1045 98.1826 49.4512C98.1826 22.75 76.0273 0.642578 49.3262 0.642578C22.6729 0.642578 0.56543 22.75 0.56543 49.4512C0.56543 76.1045 22.7207 98.2598 49.374 98.2598ZM40.2822 68.7354C37.9854 70.123 35.3535 69.0225 35.3535 66.6299V32.3203C35.3535 29.9756 38.1768 28.9707 40.2822 30.2148L68.3711 46.8672C70.3809 48.0635 70.4287 50.9346 68.3711 52.1787L40.2822 68.7354Z"
                            fill="white"/>
                    </svg>
                )}
            </div>

            <div className={styles.rightSideInfo}>
                <div className={styles.rightSideInfoContent}>
                    <span className={styles.rightSideInfoContentTitle}>good morning house mix</span>
                    <span>
                        🌴 On Vacation Hoodie → https://bit.ly/3tr3NLA
                        <br/>→ 20% off with code 'CLUB_LUNO'
                    </span>
                    <span style={{marginTop: '12px'}}>SHOW MORE</span>
                </div>

                <div className={styles.rightSideInfoChart}>
                    <WithdrawChartView price={pricePerSecond} balance={balanceLeft.current} onClick={rent}/>
                </div>
            </div>
        </div>
    </div>
}

export default Home