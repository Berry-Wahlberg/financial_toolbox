#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
量化分析模块
提供技术指标计算、回测等功能
"""

import pandas as pd
import pandas_ta as ta
import numpy as np
from typing import Dict, List, Optional


class QuantAnalyzer:
    """量化分析器"""

    def __init__(self):
        self.data = None

    def load_data(self, data: List[Dict]) -> pd.DataFrame:
        """加载数据"""
        self.data = pd.DataFrame(data)
        self.data['date'] = pd.to_datetime(self.data['date'])
        self.data.set_index('date', inplace=True)
        return self.data

    def calculate_ma(self, period: int = 20) -> pd.Series:
        """计算移动平均线"""
        if self.data is None:
            raise ValueError("请先加载数据")
        return ta.sma(self.data['close'], length=period)

    def calculate_rsi(self, period: int = 14) -> pd.Series:
        """计算RSI指标"""
        if self.data is None:
            raise ValueError("请先加载数据")
        return ta.rsi(self.data['close'], length=period)

    def calculate_macd(self, fast: int = 12, slow: int = 26, signal: int = 9) -> pd.DataFrame:
        """计算MACD指标"""
        if self.data is None:
            raise ValueError("请先加载数据")
        return ta.macd(self.data['close'], fast=fast, slow=slow, signal=signal)

    def calculate_bollinger(self, period: int = 20, std: float = 2.0) -> pd.DataFrame:
        """计算布林带"""
        if self.data is None:
            raise ValueError("请先加载数据")
        return ta.bbands(self.data['close'], length=period, std=std)


def main():
    """主函数"""
    analyzer = QuantAnalyzer()

    # 示例数据
    sample_data = [
        {'date': '2023-01-01', 'open': 100, 'high': 105, 'low': 98, 'close': 102, 'volume': 10000},
        {'date': '2023-01-02', 'open': 102, 'high': 108, 'low': 100, 'close': 106, 'volume': 12000},
    ]

    df = analyzer.load_data(sample_data)
    print("数据加载成功")
    print(df)


if __name__ == "__main__":
    main()
