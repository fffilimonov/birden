import styles from './Home.module.scss'
import classNames from 'classnames';
import img from '../assets/images/img.png'
import {useState} from 'react';
import {WithdrawChartView} from '../components/WithdrawChartVew';

const Home = () => {
    const [videoUrl, setVideoUrl] = useState();

    return <div className={styles.container}>
        <div className={classNames(styles.leftSide, styles.active)}>
            <span>Chris Luno</span>
            <span>↓ ↓ ↓</span>
            <span>YouTube</span>

            <div>
                <div>
                    <div><span>good morning house mix</span><span>birden.io/ChrisLuno/goodmorning...</span></div>
                    <div>icon 3.00</div>
                </div>
            </div>
        </div>

        <div className={styles.rightSide}>
            {!videoUrl ? <img src={img}/> : <video src={videoUrl}/>}
            <div>
                <div>
                    <span>good morning house mix</span>
                    <span>🌴 On Vacation Hoodie → https://bit.ly/3tr3NLA → 20% off with code 'CLUB_LUNO'</span>
                    <span>SHOW MORE</span>
                </div>

                <div>
                    <WithdrawChartView/>
                </div>
            </div>
        </div>
    </div>
}

export default Home