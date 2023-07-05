import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class StartedWidget extends StatelessWidget {
  const StartedWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 120,
      width: double.infinity,
      alignment: Alignment.center,
      decoration: BoxDecoration(color: Theme.of(context).colorScheme.primary),
      child: Row(
        mainAxisSize: MainAxisSize.min,
        children: [
          Padding(
            padding: const EdgeInsets.all(8.0),
            child: Text(
              AppLocalizations.of(context)!.readyToGetStarted,
              style: TextStyle(
                  fontSize: 25, color: Theme.of(context).colorScheme.onPrimary),
            ),
          ),
          const StartedButton(),
        ],
      ),
    );
  }
}

class StartedButton extends StatelessWidget {
  const StartedButton({super.key});

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
        style: ElevatedButton.styleFrom(shape: const LinearBorder()),
        onPressed: () {},
        child: Text(AppLocalizations.of(context)!.getStartedBtnText));
  }
}
