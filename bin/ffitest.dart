import 'dart:ffi';
import 'dart:isolate';

import 'package:ffi/ffi.dart';

import 'generated_bindings.dart';

void main(List<String> args) {
  final library = NativeLibrary(
    DynamicLibrary.open('ffi/lib.a'),
  );

  library.registerHandler(pointer.cast());

  library.StartServer(80);
}

typedef CFunction = Pointer<Utf8> Function(Pointer<Utf8>);
typedef DartFunction = Pointer<Utf8> Function(Pointer<Utf8>);

Pointer<Utf8> callback(Pointer<Utf8> path) {
  final pathString = path.cast<Utf8>().toDartString();
  final result = 'Hello, $pathString';
  return result.toNativeUtf8().cast<Utf8>();
}

final pointer = Pointer.fromFunction<DartFunction>(callback);
