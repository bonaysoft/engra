# engra

一本由所有学习者共同编写的英语词根词缀词典

## Goal

1. 提供词根与单词的结构化关系数据
2. 提供单词与单词的结构化关系数据
3. 提供GraphQL接口

## Non-Goal

1. 提供丰富的解释
2. 提供丰富例句
3. 提供音标朗读文件

## 数据概览

<!--START_SECTION:engra-->

|  名称   | 单词量 | 词根覆盖度 |
|---------|--------|------------|
| CEFR-A1 | 8301   | 67.97%     |
| CEFR-A2 | 11556  | 68.17%     |
| CEFR-B1 | 16691  | 73.19%     |
| CEFR-B2 | 19216  | 75.22%     |
| CEFR-C1 | 10109  | 75.37%     |
| CEFR-C2 | 9107   | 70.17%     |
| CET4    | 20855  | 76.41%     |
| CET6    | 16414  | 73.78%     |
| GMAT    | 18281  | 65.05%     |
| GRE     | 32574  | 48.35%     |
| IELTS   | 14299  | 69.40%     |
| KaoYan  | 11183  | 70.60%     |
| SAT     | 8925   | 54.08%     |
| TOEFL   | 4263   | 66.36%     |

<!--END_SECTION:engra-->

## 使用

### Graphql

```graphql endpoint doc
http://localhost:8081/
```

```graphql
{
    vocabularies(name: "weight") {
        ...vocabularyWithChildren
    }
}

fragment vocabulary on Vocabulary {
    name
}

fragment vocabularyWithChildren on Vocabulary {
    ...vocabulary
    children {
        ...vocabulary
        children {
            ...vocabulary
            children {
                ...vocabulary
                children {
                    ...vocabulary
                }
            }
        }
    }
}
```

## 迭代计划

- [x] 建立词根库
- [x] 从公开渠道爬取单词词根的数据并补充到词根库中
- [ ] 建立分级库及词频库，根据分级和词频进行优先级排序
- [ ] 按照优先级检查是否已经在词根库中，如果不再则人工补充
- [ ] 长期的人工补充工作（在前端页面查不到时引导到创建issue及PR）

## 参考

- 英语词汇的奥秘
- [外博网](https://www.waibo.wang/)
- [趣词](https://www.quword.com/)
- [ETYMONLINE](https://www.etymonline.com/)
- [ETYMOLOGEEK](https://etymologeek.com/)