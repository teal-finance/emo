package py

var TemplateStart = `from typing import Any, List, Optional


def h(msg: str):
    pass


class Emo:
    zone: Optional[str]
    activate_print = True
    activate_emojis = True
    hook = h
    _has_hook = False

    def __init__(self, *args, **kwargs):
        if len(args) == 0:
            self.zone = ""
        else:
            self.zone = args[0]
        if "activate_print" in kwargs:
            self.activate_print = kwargs["activate_print"]
        if "activate_emojis" in kwargs:
            self.activate_print = kwargs["activate_emojis"]
        if "hook" in kwargs:
            self.hook = kwargs["hook"]
            self._has_hook = True

    def _get_emo_string(self, emoji: str, obj: List[Any]) -> str:
        buf = []
        if self.activate_emojis:
            buf.append(emoji)
        if self.zone:
            buf.append("[" + self.zone + "]")
        if len(obj) > 0:
            for item in obj:
                buf.append(str(item))
        return " ".join(buf)

    def emo(self, emoji: str, obj: List[Any]) -> str:
        msg = self._get_emo_string(emoji, obj)
        if self.activate_print:
            print(msg)
        if self._has_hook:
            self.hook(msg)
        return msg

    def msg(self, *args):
        print(*args)

    def sep(self) -> str:
        msg: str = "➖➖➖➖➖➖➖➖➖➖➖"
        if self.activate_print:
            print(msg)
        return msg

    def section(self, name: str) -> str:
        msg = "➖➖➖➖➖ " + name + " ➖➖➖➖➖"
        if self.activate_print:
            print(msg)
        return msg

    def section_end(self):
        return self.sep()

    def ready(self, obj: List[Any] = []):
        if len(obj) == 0:
            obj[0] = "ready"
        return self.emo("⏲️", obj)
`