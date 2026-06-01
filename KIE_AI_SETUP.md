# KIE AI Provider Integration Guide

## Overview

Xalgorix now supports **KIE AI** as an LLM provider with full support for **Claude Opus 4.7** and other Claude models.

## Setup

### 1. Configure Environment Variables

You can set these in three ways:

#### Option A: Environment Variables (Temporary)
```bash
export XALGORIX_LLM=kie/claude-opus-4-7
export XALGORIX_API_KEY=18d1ce61db828e539211ded667b96421
export XALGORIX_API_BASE=https://api.kie.ai/claude/v1
xalgorix --target example.com --instruction "Your instruction"
```

#### Option B: Config File (Permanent)
```bash
cat > ~/.xalgorix.env << 'EOF'
XALGORIX_LLM=kie/claude-opus-4-7
XALGORIX_API_KEY=18d1ce61db828e539211ded667b96421
XALGORIX_API_BASE=https://api.kie.ai/claude/v1
EOF
```

#### Option C: Web UI Dashboard
1. Start the web UI: `xalgorix --web`
2. Navigate to Settings
3. Set the LLM Model to: `kie/claude-opus-4-7`
4. Set the API Base to: `https://api.kie.ai/claude/v1`
5. Set the API Key to: Your KIE API key

### 2. Verify Configuration

Test your setup with:
```bash
XALGORIX_DEBUG_CONFIG=1 xalgorix --target example.com --instruction "test" --max-iterations 1
```

You should see:
```
[config] Loaded: LLM="kie/claude-opus-4-7" APIBase="https://api.kie.ai/claude/v1" APIKey=****
[llm] Request → URL=https://api.kie.ai/claude/v1/messages model=claude-opus-4-7
```

## API Configuration Details

### Endpoint
- **Base URL**: `https://api.kie.ai`
- **Path**: `/claude/v1`
- **Messages Endpoint**: `/claude/v1/messages`

### Authentication
- **Method**: Bearer Token
- **Header**: `Authorization: Bearer YOUR_API_KEY`
- **API Key Format**: No prefix needed (Bearer is added automatically)

### Request Format
KIE AI uses Anthropic's message format:
```json
{
  "model": "claude-opus-4-7",
  "messages": [
    {"role": "user", "content": "Your message"}
  ],
  "system": "System prompt (optional)",
  "max_tokens": 8192,
  "stream": true/false
}
```

### Streaming Support
Both streaming and non-streaming requests are fully supported:
- **Streaming**: Server-sent events (SSE) format
- **Non-streaming**: JSON response

## Available Models

KIE AI provides access to Claude models:
- `claude-opus-4-7` (recommended)
- Other Claude models available through KIE API

Use format: `kie/model-name`

## Usage Examples

### Basic Reconnaissance
```bash
xalgorix --target example.com --instruction "Perform basic recon"
```

### Authentication Testing
```bash
xalgorix --target example.com --instruction "Test authentication mechanisms"
```

### Web Application Scanning
```bash
xalgorix --target example.com --instruction "Scan web application for vulnerabilities"
```

## Troubleshooting

### "API returned 401: Unauthorized"
- Verify API key is correct
- Check API key is active on KIE dashboard
- Ensure no typos in configuration

### "Context window exceeded"
- Claude Opus 4.7 supports 200K token context
- Xalgorix automatically prunes history to fit
- Reduce `--max-iterations` if needed

### Slow Response Times
- KIE AI usage may be rate-limited during high demand
- Default rate limit: 10 requests/second
- Adjust with `XALGORIX_RATE_RPS` if needed:
```bash
export XALGORIX_RATE_RPS=5  # Lower for stability
```

### No Response or Timeout
- Check network connectivity
- Verify API base URL is accessible
- Check API key has sufficient credits
- Increase timeout if needed:
```bash
export XALGORIX_LLM_MAX_RETRIES=5
```

## Advanced Configuration

### Rate Limiting
```bash
export XALGORIX_RATE_RPS=10        # Requests per second
export XALGORIX_RATE_BURST=20      # Burst capacity
```

### Token Management
Monitor token usage during runs. View logs with:
```bash
xalgorix --target example.com --instruction "test" 2>&1 | grep -E "\[llm\]|tokens"
```

### Custom Timeouts
```bash
export XALGORIX_LLM_MAX_RETRIES=10     # Retry attempts
export XALGORIX_MEMORY_COMPRESSOR_TIMEOUT=60  # Compressor timeout (seconds)
```

## Features

✅ **Full Streaming Support**: Stream responses in real-time
✅ **Token Accounting**: Automatic token usage tracking
✅ **Retry Logic**: Smart backoff for transient failures
✅ **Context Management**: Automatic history pruning when needed
✅ **Rate Limiting**: Configurable request throttling
✅ **Bearer Authentication**: Secure token-based auth

## Implementation Details

The following changes were made to support KIE AI:

1. **Provider Detection**: Added "kie" to recognized providers
2. **Endpoint Resolution**: Added logic to build `/messages` endpoint for KIE
3. **Authentication**: Bearer token automatically used for KIE API
4. **Request Format**: Anthropic message format (compatible with KIE)

See `internal/llm/client.go` for implementation details.

## Getting Your API Key

1. Visit [KIE AI Dashboard](https://kie.ai/api-key)
2. Create or copy your API key
3. Use it in the `XALGORIX_API_KEY` variable

## Support

For issues with KIE AI integration, check:
- [KIE AI Documentation](https://kie.ai/docs)
- [Xalgorix Issues](https://github.com/xalgord/xalgorix/issues)
- Configuration verification with `XALGORIX_DEBUG_CONFIG=1`

## Next Steps

1. ✅ Configure your KIE AI API credentials
2. ✅ Test with a simple target
3. ✅ Run full penetration testing scans
4. ✅ Monitor token usage and optimize as needed
