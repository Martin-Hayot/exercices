def capitalize(s: str):
    l = s.split(" ")
    for word in l:
        word = word[0].upper() + word[1:]
    return " ".join(l)


print(capitalize("hello world"))
