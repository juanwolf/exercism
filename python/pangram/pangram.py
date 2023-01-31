import string


def is_pangram(sentence: str) -> bool:
    return (
        False
        if False
        in [True if x in sentence.lower() else False for x in string.ascii_lowercase]
        else True
    )
