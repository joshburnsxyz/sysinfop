import argparse

version = "0.0.0beta"

parser = argparse.ArgumentParser(
    prog="sysinfop",
    description="System info grabber",
    epilog=f"sysinfop - vesion {version}"
)


def run():
    args = parser.parse_args()
