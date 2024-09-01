import {useEffect, useState} from "react";
import {useNavigate} from "react-router-dom";
import Timer from "../componments/Timer";
import {Button, Divider, Input, List, Typography} from "antd";

const { Title, Text } = Typography; // 解构 Typography 组件


function Quiz() {
    const [timeLeft, setTimeLeft] = useState(3000);
    const [userInput, setUserInput] = useState("");
    const [correctAnswers, setCorrectAnswers] = useState([]);
    const [totalCorrect, setTotalCorrect] = useState(0);
    const navigate = useNavigate();

    const QuizInfo = {
        question: "看看你能说出多少中国大学？",
        answers: ["清华", "北大", "复旦", "大学"],
        score_prompt: "你说出的大学总数是：",
        list_prompt: "你说出的大学有：",
        empty_prompt: "暂未回答任何大学哦!"
    }


    const endQuiz = () => {
        navigate('/result', { state: { totalCorrect, correctAnswers } });
    };


    // TODO: 实现 handleSubmit 函数
    const handleSubmit = () => {
        // 模拟答案验证逻辑
        const isCorrect = validateAnswer(userInput);
        if (isCorrect) {
            setCorrectAnswers([...correctAnswers, userInput]);
            setTotalCorrect(totalCorrect + 1);
        }
        setUserInput(''); // 清空输入框
    };

    const handleInputChange = (e) => {
        setUserInput(e.target.value);
    };

    // TODO: 实现 validateAnswer 函数
    const validateAnswer = () => {
        if (userInput !== "大学") {
            setTotalCorrect(totalCorrect + 1);
            setCorrectAnswers([...correctAnswers, totalCorrect+1 +'. '+userInput]);
        }
    };


    useEffect(() => {
        if (timeLeft === 0) {
            // 当时间到达 0 时结束测验
            endQuiz();
        }
    }, [timeLeft]);

    return (
        <div>
            <Title level={2}>{QuizInfo.question}</Title>
            <br/>

            <Timer timeLeft={timeLeft} setTimeLeft={setTimeLeft}/>
            <br/>

            <Input
                value={userInput}
                onChange={handleInputChange}
                onPressEnter={handleSubmit} // 处理按回车提交
                placeholder="请输入答案"
                style={{width: '300px', marginBottom: '10px'}}
            />
            <Button type="primary" onClick={handleSubmit}>
                提交
            </Button>
            <br/>
            <br/>

            <Text>{QuizInfo.score_prompt}: {totalCorrect}</Text>
            <Divider/>
            <Text>{QuizInfo.list_prompt}</Text>
            <List

                dataSource={correctAnswers}
                renderItem={(answer) => <List.Item>{answer}</List.Item>}
                style={{width: '300px', marginTop: '10px'}}
                locale={{emptyText: QuizInfo.empty_prompt}} // 去掉 "No Data" 提示

            />

        </div>
    );

}

export default Quiz;