import React, { useEffect } from 'react';

const Timer = ({ timeLeft, setTimeLeft }) => {
    useEffect(() => {
        if (timeLeft > 0) {
            const timerId = setInterval(() => {
                setTimeLeft((prevTime) => prevTime - 1);
            }, 1000);

            return () => clearInterval(timerId); // 清除定时器
        }
    }, [timeLeft, setTimeLeft]);

    const formatTime = (time) => {
        const minutes = Math.floor(time / 60);
        const seconds = time % 60;
        return `${minutes}:${seconds < 10 ? `0${seconds}` : seconds}`;
    };

    return <div>剩余时间: {formatTime(timeLeft)}</div>;
};

export default Timer;
