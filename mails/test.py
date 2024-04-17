import requests
from dataclasses import asdict
from dataclasses import dataclass
from typing import Optional
import json
import pandas as pd 


@dataclass
class Request:
    category_id: Optional[str] = None
    is_activity: Optional[str] = None
    is_cjmj: Optional[str] = None
    is_dpby: Optional[str] = None
    is_jxq: Optional[str] = None
    is_mzhg: Optional[str] = None
    is_xy: Optional[str] = None
    is_youx: Optional[str] = None
    no_image: Optional[str] = None
    order: Optional[str] = None
    page: Optional[str] = None
    sort: Optional[str] = None


# 创建一个Request对象
params = Request(
    # category_id="<category_id>",
    # is_activity="<is_activity>",
    # is_cjmj="<is_cjmj>",
    # is_dpby="<is_dpby>",
    # is_jxq="<is_jxq>",
    # is_mzhg="<is_mzhg>",
    # is_xy="<is_xy>",
    # is_youx="<is_youx>",
    # no_image="<no_image>",
    # order="<order>",
    page="<page>",
    # sort="<sort>"
)


class CollyMedicine:
    def __init__(self) -> None:
        self.headers = {
        'Cookie': 'longyiyy_pc_session=eyJpdiI6IjMva3VXV2VjbkxKMmMydzljeGxTNFE9PSIsInZhbHVlIjoiY21pb014QTBEZ21uaW8vdEpGekUzaWd0U0U5QUdkUDZYaXhmNy8vd21ZOWJRdXpNZkVVZWxIaEhKTkhxbFNMd2kwbEQvR0dQVU96TXhXZ3hzSDdJODV5NjFFNXhSK2E3TGw3MGlBM3ZQVVN3ZWNzV3dmem56d3RLZzVZSFk3UUkiLCJtYWMiOiJkNzg2NDRlN2E5YjMxMWEyNGIyMWI4NjdiOTdjMjJmMTQwNGUwNjc3MzJhOWJiODk2MGU5ZjAyZjZjMjA3ZDg2In0%3D',
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36'
        }
        self.url = "https://pc.api.longyiyy.com/api/goods/list"

    @staticmethod
    def append_to_csv(df:pd.DataFrame):
        csv_file = "colly.csv"
        try:
            with open(csv_file, 'x') as f:
                df.to_csv(f, index=False)
        except FileExistsError:
            # 如果文件已存在，则以追加模式写入数据，不再包含列名
            with open(csv_file, 'a') as f:
                df.to_csv(f, index=False, header=False)




    def parseBody(self,body,page):
        csvData = []
        fmtBody = json.loads(body)
        for item in fmtBody['data']:
            # print(f"name:{item['name']}   price:{item['price']}")
            csvData.append(
                {'name': item['name']
                 ,'price': item['price']
                 # 这里自由发挥，将哪些写入csv啥的
                 }
                )
        print(f"保存第 {page} 页数据ing...")
        df = pd.DataFrame(csvData)
        self.append_to_csv(df)

    def run(self):
        # 控制页数
        for page in range(1,5):
            params = Request(page=page)
            # 将Request对象转换为字典
            params_dict = asdict(params)
            response = requests.get(self.url, headers=self.headers, params=params_dict)
            # print(response.text)
            self.parseBody(response.text,page)
        print(f"保存完毕，请查看本地colly.csv文件")


if __name__ == "__main__":
    colly = CollyMedicine()
    colly.run()
    




