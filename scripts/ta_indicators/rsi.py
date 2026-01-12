#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
RSI指标计算脚本
功能：读取stdin中的K线数据，计算RSI指标，输出结果到stdout
"""

import sys
import json
import pandas as pd
import pandas_ta as ta

def main():
    """主函数"""
    try:
        # 从stdin读取JSON数据
        input_data = sys.stdin.read()
        if not input_data:
            raise ValueError("输入数据为空")
        
        # 解析JSON数据
        kline_data = json.loads(input_data)
        
        # 转换为DataFrame
        df = pd.DataFrame(kline_data)
        
        # 检查必要的字段
        required_fields = ['open', 'high', 'low', 'close', 'volume', 'date']
        for field in required_fields:
            if field not in df.columns:
                raise ValueError(f"缺少必要字段: {field}")
        
        # 计算RSI指标（默认周期14）
        rsi_results = ta.rsi(df['close'])
        
        # 如果结果为空，抛出错误
        if rsi_results is None:
            raise ValueError("RSI计算失败，结果为空")
        
        # 合并原始日期和RSI结果
        result_df = pd.concat([df[['date']], rsi_results], axis=1)
        
        # 将结果转换为JSON并输出到stdout
        result_json = result_df.to_json(orient='records', force_ascii=False)
        sys.stdout.write(result_json)
        sys.stdout.flush()
        
    except Exception as e:
        # 捕获所有异常，输出错误信息到stderr
        error_response = {
            "success": False,
            "error": str(e)
        }
        error_json = json.dumps(error_response, ensure_ascii=False)
        sys.stderr.write(error_json)
        sys.stderr.flush()
        sys.exit(1)

if __name__ == "__main__":
    main()