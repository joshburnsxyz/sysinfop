from rich.layout import Layout
from rich.panel import Panel
from .sysinfo import getSystemInfo

layout = Layout()
data = getSystemInfo()
banner = Layout(name=f"{data['hostname']}")
lower = Layout(name="lower")

layout.split_column(
    banner,
    lower
)

layout["lower"].split_row(
    Layout(name="left"),
    Layout(name="right"),
)
