BOTY_PACKAGE_PATH := ./apps/boty

build-boty-package:
	cd $(BOTY_PACKAGE_PATH) && npm run bootstrap && npm run build

run_boty_package:
	cd $(BOTY_PACKAGE_PATH) && npm run start
