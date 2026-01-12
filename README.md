# Personal Financial Toolbox

## Purpose

The Personal Financial Toolbox is a comprehensive application designed to provide advanced financial analysis capabilities. It leverages the power of Go for backend API development and Python for complex data analysis, enabling users to perform technical analysis on stock data, including indicators like MACD, RSI, and backtesting strategies.

## Architecture

- **Backend**: Go 1.21+ (Controller and Service layers)
- **Data Analysis**: Python 3.10+ (Worker scripts for complex calculations)
- **Communication**: JSON data exchange between Go and Python processes
- **API Framework**: Gin Web Framework

## Installation

### Prerequisites

- Go 1.21 or higher
- Python 3.10 or higher
- Git

### Steps

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd financial_toolbox
   ```

2. **Install Go dependencies**
   ```bash
   go mod download
   ```

3. **Install Python dependencies**
   ```bash
   pip install -r scripts/requirements.txt
   ```

## Usage

### Running the Application

1. **Start the Go server**
   ```bash
   go run main.go
   ```

2. **Access the API**
   The server will start on `http://localhost:8080`

### API Endpoints

#### GET /api/indicator

Calculates technical indicators for a given stock symbol.

**Parameters**:
- `symbol`: Stock symbol (e.g., "AAPL", "GOOGL")
- `indicator`: Indicator type (e.g., "macd")

**Example Request**:
```bash
curl "http://localhost:8080/api/indicator?symbol=AAPL&indicator=macd"
```

**Example Response**:
```json
{
  "success": true,
  "result": "[{\"date\":\"2023-01-01\",\"MACD_12_26_9\":null,\"MACDh_12_26_9\":null,\"MACDs_12_26_9\":null},{\"date\":\"2023-01-02\",\"MACD_12_26_9\":null,\"MACDh_12_26_9\":null,\"MACDs_12_26_9\":null},...]"
}
```

## Features

- **MACD Indicator**: Moving Average Convergence Divergence calculation
- **RESTful API**: Easy-to-use API for accessing financial indicators
- **Modular Architecture**: Separation between backend and data analysis layers
- **Error Handling**: Comprehensive error handling for robust operation
- **Timeout Protection**: Prevents hanging from unresponsive Python scripts

## Planned Features

- **Additional Indicators**: RSI, Bollinger Bands, SMA, EMA, etc.
- **Backtesting Framework**: Test trading strategies against historical data
- **Real-time Data Integration**: Connect to financial data APIs for live data
- **Visualization**: Generate charts and graphs for indicator results
- **User Authentication**: Secure access to the API

## Contribution

### Branching Strategy

This project uses a simplified Git Flow branching strategy combined with trunk-based development:

1. **Main Branch**: Always represents production-ready code
2. **Feature Branches**: Created for new development work
   - Naming convention: `feature/feature-name`
   - Short-lived branches (usually 1-3 days)
3. **Code Review**: All feature branches require review before merging to main
4. **Testing**: Ensure all changes are tested before integration

### Contribution Workflow

1. **Create a feature branch**
   ```bash
   git checkout main
   git pull origin main
   git checkout -b feature/your-feature-name
   ```

2. **Make changes** and **commit** them
   ```bash
   git add .
   git commit -m "Description of your changes"
   ```

3. **Push** the branch to the repository
   ```bash
   git push origin feature/your-feature-name
   ```

4. **Create a Pull Request** for code review

5. **Merge** the branch after approval and successful tests

### Code Guidelines

- **Go Code**: Follow `gofmt` standards, include Chinese comments
- **Python Code**: Follow PEP 8 style guide
- **Documentation**: Update documentation for new features
- **Tests**: Add tests for new functionality

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For questions or suggestions, please open an issue on the repository.