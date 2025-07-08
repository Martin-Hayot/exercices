def join(sep, *args):
    if not args:
        return ""
    result = str(args[0])
    for arg in args[1:]:
        result += sep + str(arg)
    return result


if __name__ == "__main__":
    # Test cases
    j1 = join(":", "Bonjour", "tout", "le", "monde")
    print(f"join(':',  'Bonjour', 'tout', 'le', 'monde') = '{j1}'")

    j2 = join("-", "Hello", "World")
    print(f"join('-', 'Hello', 'World') = '{j2}'")

    j3 = join(",", "1", "2", "3")
    print(f"join(',', '1', '2', '3') = '{j3}'")

    # Edge cases
    j4 = join(":", "single")
    print(f"join(':', 'single') = '{j4}'")

    j5 = join(":")
    print(f"join(':') = '{j5}'")

    # Mixed types (numbers, etc.)
    j6 = join("-", "item", 1, 2.5, True)
    print(f"join('-', 'item', 1, 2.5, True) = '{j6}'")
