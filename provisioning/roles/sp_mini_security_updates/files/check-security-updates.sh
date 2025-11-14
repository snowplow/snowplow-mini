#!/bin/bash
# Script to check for available security updates on Snowplow Mini
# This can be run manually to see what updates are pending

set -e

echo "============================================"
echo "Snowplow Mini - Security Update Check"
echo "============================================"
echo ""

# Check if running as root
if [ "$EUID" -ne 0 ]; then
    echo "Please run as root (use sudo)"
    exit 1
fi

# Update package lists
echo "[1/4] Updating package lists..."
apt-get update -qq

# Check for security updates
echo ""
echo "[2/4] Checking for security updates..."
SECURITY_UPDATES=$(apt-get upgrade -s | grep -i security | wc -l)

if [ "$SECURITY_UPDATES" -gt 0 ]; then
    echo "⚠️  Found $SECURITY_UPDATES security updates available"
    echo ""
    apt-get upgrade -s | grep -i security
else
    echo "✅ No security updates available"
fi

# Check for all updates (informational only)
echo ""
echo "[3/4] Checking for all available updates..."
apt list --upgradable 2>/dev/null | grep -v "Listing..." | head -20

# Show unattended-upgrades status
echo ""
echo "[4/4] Unattended-upgrades status..."
systemctl status apt-daily.timer --no-pager | head -10
systemctl status apt-daily-upgrade.timer --no-pager | head -10

# Show recent logs
echo ""
echo "Recent unattended-upgrades logs:"
if [ -f /var/log/unattended-upgrades/unattended-upgrades.log ]; then
    tail -20 /var/log/unattended-upgrades/unattended-upgrades.log
else
    echo "(No logs yet - updates haven't run)"
fi

echo ""
echo "============================================"
echo "To manually run security updates:"
echo "  sudo unattended-upgrade -d"
echo ""
echo "To disable automatic updates:"
echo "  Edit /etc/snowplow-mini-security.conf"
echo "============================================"
