# CI/CD Pipeline Documentation

## Overview

This document describes the automated build and deployment pipeline for Financial Toolbox using GitHub Actions.

## Workflow Triggers

The CI/CD pipeline is triggered on:

- **Push** to `main` or `develop` branches
- **Pull requests** to `main` branch
- **Manual dispatch** via GitHub Actions UI

## Build Jobs

### 1. Windows Build (amd64)
- **Runner**: `windows-latest`
- **Platform**: Windows/amd64
- **Output**: `financial-toolbox.exe`
- **Dependencies**: Go 1.21, Node.js 18, Wails CLI

### 2. Linux Build (amd64)
- **Runner**: `ubuntu-latest`
- **Platform**: Linux/amd64
- **Output**: `financial-toolbox` (executable)
- **Dependencies**: Go 1.21, Node.js 18, Wails CLI, GTK3, WebKit2GTK

### 3. macOS Build (amd64)
- **Runner**: `macos-latest`
- **Platform**: Darwin/amd64
- **Output**: `Financial Toolbox.app`
- **Dependencies**: Go 1.21, Node.js 18, Wails CLI

### 4. Release Job
- **Condition**: Only runs on push to `main` branch
- **Creates**: GitHub Release with all platform artifacts
- **Version**: `v{run_number}` (e.g., v123)

## Caching Strategy

The workflow uses caching to speed up builds:

- **Go modules**: Cached based on `go.sum` hash
- **Node modules**: Cached based on `package-lock.json` hash
- **Cache retention**: Follows GitHub's default cache retention policy

## Artifacts

Build artifacts are stored for 30 days and include:

- `financial-toolbox-windows-amd64`: Windows executable
- `financial-toolbox-linux-amd64`: Linux executable
- `financial-toolbox-macos-amd64`: macOS application bundle

## Manual Build

### Windows
```bash
wails build -clean -platform windows/amd64
```

### Linux
```bash
sudo apt-get update
sudo apt-get install -y libgtk-3-dev libwebkit2gtk-4.0-dev
wails build -clean -platform linux/amd64
```

### macOS
```bash
wails build -clean -platform darwin/amd64
```

## Troubleshooting

### Build Failures

1. **Check logs**: Review the GitHub Actions logs for specific error messages
2. **Dependencies**: Ensure all system dependencies are installed
3. **Version compatibility**: Verify Go and Node.js versions match the workflow

### Cache Issues

If caching causes problems, you can clear caches:
1. Go to repository Settings → Actions → Caches
2. Delete the problematic cache entries
3. Re-run the workflow

### Release Issues

If the release job fails:
1. Verify the tag doesn't already exist
2. Check that `GITHUB_TOKEN` has proper permissions
3. Ensure artifact paths are correct

## Local Testing

To test the workflow locally:

1. Install [act](https://github.com/nektos/act):
   ```bash
   # macOS/Linux
   brew install act

   # Windows
   choco install act-cli
   ```

2. Run the workflow:
   ```bash
   act push
   ```

## Security Considerations

- **Secrets**: No custom secrets required (uses default `GITHUB_TOKEN`)
- **Permissions**: Workflow has read/write permissions for releases
- **Dependencies**: All dependencies are from official sources

## Future Improvements

- [ ] Add code signing for Windows/macOS
- [ ] Implement automated testing
- [ ] Add Docker container builds
- [ ] Set up automated versioning
- [ ] Add build notifications
- [ ] Implement rollback mechanism

## Support

For issues or questions:
- Open an issue on GitHub
- Check the [Wails documentation](https://wails.io/docs/)
- Review [GitHub Actions documentation](https://docs.github.com/en/actions)
