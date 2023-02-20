# engra

一本由所有学习者共同编写的英语词根词缀词典

## 数据概览

<!--START_SECTION:engra-->

|  名称   | 单词量 | 词根覆盖度 |
|---------|--------|------------|
| CEFR-A1 | 592    | 67.23%     |
| CEFR-A2 | 888    | 67.45%     |
| CEFR-B1 | 1390   | 72.66%     |
| CEFR-B2 | 1746   | 74.34%     |
| CEFR-C1 | 1010   | 73.07%     |
| CEFR-C2 | 1011   | 68.15%     |
| CET4    | 2607   | 74.95%     |
| CET6    | 2345   | 71.86%     |
| GMAT    | 3047   | 62.42%     |
| GRE     | 6515   | 45.22%     |
| IELTS   | 3575   | 67.41%     |
| KaoYa   | 3728   | 68.86%     |
| SAT     | 4463   | 49.09%     |
| TOEFL   | 4264   | 63.95%     |

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

## 前端

提供接口和组件两种形式，组件默认展示在官方网站上。

形式为：搜索框，搜索任意单词，出现该次的构词及词根树

## 数据

数据以词根为单元采用yaml文件存储在dicts中。首批词根来源于《英语词汇的奥秘》升级版

## 迭代计划

- [x] 建立词根库
- [x] 从公开渠道爬取单词词根的数据并补充到词根库中
- [ ] 建立分级库及词频库，根据分级和词频进行优先级排序
- [ ] 按照优先级检查是否已经在词根库中，如果不再则人工补充
- [ ] 长期的人工补充工作（在前端页面查不到时引导到创建issue及PR）