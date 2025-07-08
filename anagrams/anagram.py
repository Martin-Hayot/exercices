def anagram(s1, s2) -> bool:
    if s1 == "" or s2 == "":
        return False
    if sorted(s1) == sorted(s2):
        return True
    return False


print(anagram("hello", "elloh"))
