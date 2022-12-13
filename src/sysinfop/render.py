from rich.layout import Layout
from rich.panel import Panel
from .sysinfo import getSystemInfo

layout = Layout()
data = getSystemInfo()

layout.split_column(
    Layout(name=f"{data.hostname}"),
    Layout(name="lower")
)

layout["lower"].split_row(
    Layout(name="left"),
    Layout(name="right"),
)

layout["right"].split(
    Layout(Panel("Hello")),
    Layout(Panel("World!"))
)

layout["left"].update(
    "The mystery of life isn't a problem to solve, but a reality to experience."
)
