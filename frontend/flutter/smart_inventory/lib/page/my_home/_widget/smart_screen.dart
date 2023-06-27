import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:smart_inventory/page/my_home/_widget/screen/purchase_order_list.dart';
import 'package:smart_inventory/provider/error_refresh.dart';
import 'package:smart_inventory/provider/smart_provider.dart';

class SmartScreenWidget extends ConsumerStatefulWidget {
  const SmartScreenWidget({super.key});

  @override
  ConsumerState<ConsumerStatefulWidget> createState() =>
      _SmartScreenWidgetState();
}

class _SmartScreenWidgetState extends ConsumerState<SmartScreenWidget> {
  @override
  Widget build(BuildContext context) {
    final smartHome = ref.watch(smartHomeEnquireProvider);
    switch (smartHome.currentModel) {
      case "ListPurchaseOrders":
        return const PurchaseOrderListScreen();
      default:
        return Container();
    }
  }
}
