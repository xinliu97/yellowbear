import {useEffect, useState} from "react";
import {useNavigate} from "react-router-dom";
import Timer from "../componments/Timer";
import {Button, Divider, Input, Typography, List, Row, Col, Progress} from "antd";
import axios from 'axios'
import InfiniteScroll from 'react-infinite-scroll-component';
const { Title, Text } = Typography; // 解构 Typography 组件


function Quiz() {
    const [timeLeft, setTimeLeft] = useState(1);
    const [userInput, setUserInput] = useState("");
    const [correctAnswers, setCorrectAnswers] = useState([]);
    const [totalCorrect, setTotalCorrect] = useState(0);
    const [isQuizEnded, setIsQuizEnded] = useState(false);
    const [resultStatsData, setResultStatsData] = useState(null);
    const [resultStatsTitle, setResultStatsTitle] = useState(null);
    const navigate = useNavigate();

    const QuizInfo = {
        quiz_id:12345,
        question: "看看你能说出多少中国大学？",
        answers: ["清华", "北大", "复旦", "大学"],
        score_prompt: "你说出的大学总数是：",
        list_prompt: "你说出的大学有：",
        empty_prompt: "暂未回答任何大学哦!"
    }


    const endQuiz = () => {
        sendResultsToBackend();


    };
    const sendResultsToBackend = () =>{
        const data = {
            quiz_id: QuizInfo.quiz_id,
            total_correct: totalCorrect,
            correct_answers: correctAnswers,
            user_id:1234
        }

        axios.post('http://127.0.0.1:5000/handle_submit', data)
            .then((response) => {
                    console.log("Send successfully!");
                    setResultStatsTitle(response.data.title);

                    setResultStatsData(response.data.choices);

                    //correct data
                    console.log(response.data);
                    //null
                    console.log(resultStatsData)

                })
            .catch((error) => {
                console.error("Error when sending result!", error)

            })

        // use fetch to send a post request
    }


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
        if (resultStatsData !== null) {
            console.log("Updated resultStatsData:", resultStatsData);
            console.log("Updated resultStatsTitle:", resultStatsTitle);
            setIsQuizEnded(true);
        }
    }, [resultStatsData]);

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

            {isQuizEnded ? (
                    <div>
                        <Text>{QuizInfo.score_prompt}: {totalCorrect}</Text>
                        <Divider />
                        <Text>今天答题人数:</Text>
                        <List
                            grid={{ gutter: 16, xs: 1, sm: 2, md: 2, lg: 2, xl: 2 }} // 自适应布局，控制每行最多展示两个元素
                            dataSource={resultStatsData}
                            header={
                                <Row style={{ width: '100%', fontWeight: 'bold' }}>
                                    <Col flex="100px">
                                        <Text>{resultStatsTitle}</Text>
                                    </Col>
                                    <Col flex="auto">
                                        <Text>Correct Percentage</Text>
                                    </Col>
                                </Row>
                            }
                            renderItem={(item, index) => (
                                <List.Item>

                                    <Row align="middle" style={{ width: '100%' }}>
                                        {/* 标号和Key */}
                                        <Col flex="100px">
                                            <Text>{index + 1}. {item.choice_key}</Text>
                                        </Col>
                                        {/* 百分比进度条 */}
                                        <Col flex="auto">
                                            <Progress percent={item.choice_value} status="active" />
                                        </Col>
                                        {/* 数值 */}
                                        {/*<Col flex="50px">*/}
                                        {/*    <Text>{item.choice_value}%</Text>*/}
                                        {/*</Col>*/}
                                    </Row>
                                </List.Item>
                            )}
                        />

                    </div>

                ):(
                    <div>
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
            )}
        </div>
    );

}

export default Quiz;