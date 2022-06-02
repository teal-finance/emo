import 'package:emodebug/emodebug.dart';

/// The normal [EmoDebug] class initialization
//const EmoDebug debug = EmoDebug();

/// The custom implementation
const CustomEmoDebug debug = CustomEmoDebug();

/// A class that extends [EmoDebug] with custom methods
class CustomEmoDebug extends EmoDebug {
  /// The base constructor
  const CustomEmoDebug();

  /// The crash method to indicate a program failure
  void crash(dynamic obj, [String? domain]) => emo("💥", obj, domain);

  /// The recovery method to indicate a program recovery
  void recovery(dynamic obj, [String? domain]) => emo("👍", obj, domain);
}

void main() {
  _printLines(1);
  debug.init("Initializing");
  _printLines(3);
  debug.ok("Everything is ok");
  _printLines(2);
  debug.state("A state operation");
  _printLines(1);
  debug.save("Saving something");
  _printLines(2);
  debug.delete("Deleting something");
  _printLines(3);
  debug.update("Updating something");
  _printLines(2);
  _printLines(1);
  final data = {"foo": "bar"};
  debug.data(data, "some data");
  _printLines(2);
  debug.crash("A crash occured!!!");
  _printLines(4);
  debug.recovery("Recovery successful");
  _printLines(3);
  debug.emo("🏁", "Finish");
}

var _i = 1;

void _printLines(int n) {
  var i = 0;
  while (i < n) {
    print("[$_i] Normal debug message flow message");
    ++i;
    ++_i;
  }
}
