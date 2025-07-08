

def isDigit(s):
    if not isinstance(s, str):
        return False

    # Handle empty string
    if not s:
        return False
    try:
        if isinstance(int(s), int):
            return True
    except:
        return False


if __name__ == "__main__":
    number = "-112112"

    print(isDigit(number))
