import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:smart_inventory/provider/error_refresh.dart';
import 'package:smart_inventory/provider/smart_provider.dart';
import 'package:smart_inventory/src/generated/smart/v1/smart.pb.dart';

class SlideInInputWidget extends ConsumerStatefulWidget {
  const SlideInInputWidget({super.key});

  @override
  ConsumerState<SlideInInputWidget> createState() => SlideInInputWidgetState();
}

class SlideInInputWidgetState extends ConsumerState<SlideInInputWidget>
    with SingleTickerProviderStateMixin {
  late final AnimationController _animationController;
  late final Animation<double> _animation;

  bool _isFold = true;

  @override
  void initState() {
    super.initState();
    _animationController = AnimationController(
      vsync: this,
      duration: const Duration(milliseconds: 300),
    );
    _animation = Tween<double>(begin: 1.0, end: 0.0).animate(CurvedAnimation(
      parent: _animationController,
      curve: Curves.easeIn,
    ));
  }

  void startAnimation() async {
    if (_animationController.isAnimating) return;
    if (_isFold) {
      _animationController.forward();
      _isFold = false;
    } else {
      _animationController.reverse();
      _isFold = true;
    }
  }

  void foldAnimation() {
    if (_isFold) {
      return;
    }
    startAnimation();
  }

  @override
  void dispose() {
    _animationController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return AnimatedBuilder(
      animation: _animation,
      builder: (BuildContext context, Widget? child) {
        return Transform.scale(
          scaleX: _animation.value,
          child: Opacity(
            opacity: _animation.value,
            child: const SmartChatWidget(),
          ),
        );
      },
    );
  }
}

class SmartChatWidget extends ConsumerStatefulWidget {
  const SmartChatWidget({super.key});

  @override
  ConsumerState<ConsumerStatefulWidget> createState() =>
      _SmartChatWidgetState();
}

class _SmartChatWidgetState extends ConsumerState<SmartChatWidget> {
  final TextEditingController _inputController = TextEditingController();
  SmartChat? _smartChat;

  @override
  void dispose() {
    _inputController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final liveSmartChat = ref.watch(smartHomeEnquireSmartChatProvider);
    return Column(
      mainAxisSize: MainAxisSize.max,
      children: [
        liveSmartChat.when(data: (sc) {
          _smartChat = sc;
          return ListView.builder(
              shrinkWrap: true,
              itemCount: sc.messages.length,
              padding: const EdgeInsets.only(right: 80),
              itemBuilder: (context, index) {
                final message = sc.messages[index];
                return _ChatHistoryMessageWidget(
                  message.content,
                  sender: message.sender,
                );
              });
        }, error: (error, _) {
          // ref.read(errorRefreshProvider.notifier).markAsNeedRefresh();
          return Container();
        }, loading: () {
          if (_smartChat == null) {
            return Container();
          }
          return const CircularProgressIndicator();
        }),
        const SizedBox(
          height: 3,
        ),
        Padding(
          padding: const EdgeInsets.symmetric(horizontal: 8.0),
          child: TextFormField(
            controller: _inputController,
            decoration: const InputDecoration(
              labelText: 'Input',
              enabledBorder: OutlineInputBorder(
                borderRadius: BorderRadius.all(Radius.circular(10)),
              ),
              focusedBorder: OutlineInputBorder(
                borderRadius: BorderRadius.all(Radius.circular(10)),
              ),
              errorBorder: OutlineInputBorder(
                borderRadius: BorderRadius.all(Radius.circular(10)),
              ),
              focusedErrorBorder: OutlineInputBorder(
                borderRadius: BorderRadius.all(Radius.circular(10)),
              ),
            ),
            onFieldSubmitted: (v) {
              ref
                  .read(smartHomeEnquireProvider.notifier)
                  .sender
                  .add(_smartChat ?? SmartChat()
                    ..messages.add(ChatMessage()..content = v));

              _inputController.text = '';
            },
          ),
        ),
      ],
    );
  }
}

class _ChatHistoryMessageWidget extends StatelessWidget {
  final String text;
  final String sender;
  const _ChatHistoryMessageWidget(this.text, {required this.sender});

  @override
  Widget build(BuildContext context) {
    return Container(
      alignment: sender == 'GPT' ? Alignment.centerLeft : Alignment.centerRight,
      width: double.infinity,
      child: Text(text),
    );
  }
}
