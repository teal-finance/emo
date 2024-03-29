package dart

var codeStart = `// Code generated by https://github.com/teal-finance/emo/blob/main/codegen/dart/gen.go ; DO NOT EDIT.

import 'package:meta/meta.dart';
import 'types.dart';

/// The debug printer
@immutable
class EmoDebug {
  /// Default constructor
  const EmoDebug(
      {this.zone,
      this.hook,
      this.deactivatePrint = false,
      this.deactivateEmojis = false});

  /// Deactivate the debug session
  ///
  /// This will not print anything anymore
  final bool deactivatePrint;

  /// The debug zone
  ///
  /// An optional indication about a local debug area. It
  /// will prefix the messages
  final String? zone;

  /// A hook to execute after each function call
  ///
  /// Can be used to pipe the emodebug messages to logging
  final EmoDebugHook? hook;

  /// Deactivate the emojis
  ///
  /// This will not print the emojis in the log messages
  final bool deactivateEmojis;

  /// A debug message for a ready state
  ///
  /// emoji: ⏲️
  String ready([dynamic obj, String? domain]) {
    obj ??= "ready";
    return emo("⏲️", obj, domain);
  }

  /// A simple message with no emoji
  String msg([dynamic obj, String? domain]) => emo(null, obj, domain);

  /// Print a debug message from an emoji
  String emo(String? emoji, [dynamic obj, String? domain]) {
    final msg = _getEmoString(emoji, obj, domain);
    if (!deactivatePrint) {
      print(msg);
    }
    hook?.call(msg);
    return msg;
  }

  /// A separator line
  String sep() {
    const msg = "➖➖➖➖➖➖➖➖➖➖➖";
    if (!deactivatePrint) {
      print(msg);
    }
    hook?.call(msg);
    return msg;
  }

  /// A section start
  String section(String name) {
    final msg = "➖➖➖➖➖ $name ➖➖➖➖➖";
    if (!deactivatePrint) {
      print(msg);
    }
    hook?.call(msg);
    return msg;
  }

  /// A section end
  String sectionEnd() => sep();

  String _getEmoString(String? emoji, dynamic obj, String? domain) {
    final l = <String>[];
    if (!deactivateEmojis && emoji != null) {
      l.add("$emoji");
    }
    if (zone != null) {
      l.add("[$zone]");
    }
    if (domain != null) {
      final dm = '${domain[0].toUpperCase()}${domain.substring(1)}:';
      l.add(dm);
    }
    if (obj != null) {
      l.add("$obj");
    }
    //print("$obj ======= $emoji");
    //print(l);
    return l.join(" ");
  }
`

var codeEnd = `}`
