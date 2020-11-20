#!/usr/bin/env bash
# bash

sh tdd/clean-testcache.sh

sh tdd/ut.db.migration.run-001.sh

sh tdd/ut.pkg.infra-layer.repository-001.sh

sh tdd/ut.pkg.app-layer.service-001.sh

sh tdd/ut.pkg.app-layer.application-001.sh

# sh tdd/ut.db.migration.rollback-001.sh