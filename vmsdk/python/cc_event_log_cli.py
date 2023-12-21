"""
Command line to dump the cc event logs
"""
import logging
import argparse
from cctrusted import CCTrustedVmSdk


LOG = logging.getLogger(__name__)
OUT_FORMAT_RAW = "raw"
OUT_FORMAT_HUMAN = "human"

logging.basicConfig(level=logging.NOTSET, format='%(message)s')

def out_format_validator(out_format):
    """Validator (callback for ArgumentParser) of output format
    Args:
        out_format: User specified output format.
    Returns:
        Validated value of the argument.
    Raises:
        ValueError: An invalid value is given by user.
    """
    if out_format not in (OUT_FORMAT_HUMAN, OUT_FORMAT_RAW):
        raise ValueError
    return out_format

def main():
    """example cc event log fetching utility"""
    parser = argparse.ArgumentParser(
        description="The example utility to fetch CC event logs")
    parser.add_argument('-s', type=int, default=0,
                        help='index of first event log to fetch', dest='start')
    parser.add_argument("-c", type=int, help="number of event logs to fetch",
                        dest="count")
    parser.add_argument("--out-format", default=OUT_FORMAT_RAW,
                        dest="out_format",
                        help="Output format: raw/human. Default raw.",
                        type=out_format_validator)
    args = parser.parse_args()

    event_logs = CCTrustedVmSdk.inst().get_eventlog(args.start, args.count)
    if event_logs is not None:
        LOG.info("Total %d of event logs fetched.", len(event_logs.event_logs))
        event_logs.dump(args.out_format == OUT_FORMAT_RAW)

if __name__ == "__main__":
    main()