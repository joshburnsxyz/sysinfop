from rich.layout import Layout
from rich.panel import Panel
from rich import box
from .sysinfo import getSystemInfo

layout = Layout()
data = getSystemInfo()
banner = Layout(name=f"{data['hostname']}",ratio=4)
logo = Layout(name=f"{data['platform']}",ratio=1)
box = box.MINIMAL

layout.split_row(
    logo,
    banner
)

# System info section
h = data["hostname"]
i = data["ip-address"]
p = data["platform"]
P = data["processor"]
r = data["ram"]
banner.update(Panel(f"* {h}\n* IP: {i}\n* OS: {p}\n* CPU: {P}\n* RAM: {r}", title=data["hostname"], expand=False, height=10, box=box))

# Logo Section
logo.update(Panel(f"{data['logo']}", expand=False, title=data["platform"], height=10, box=box))
