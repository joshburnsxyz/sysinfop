from rich.layout import Layout
from rich.panel import Panel
from .sysinfo import getSystemInfo

layout = Layout()
data = getSystemInfo()
banner = Layout(name=f"{data['hostname']}")

layout.split_column(
    banner
)

# Top banner, basic info
banner.update(Panel(f"* {data["hostname"]}\n* IP: {data["ip-address"]}\n* OS: {data["platform"]}\n* CPU: {data["processor"]}", title=data["hostname"])
