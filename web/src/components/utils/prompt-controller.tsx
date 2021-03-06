/**
 * Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import React from 'react';
import { gql, useQuery } from '@apollo/client';
import { Organization } from 'Generated/schema';
import useModal from 'Hooks/useModal';
import { MODALS } from 'Components/utils/modal-context';

// We are intentionally over-fetching, in order to proactively add this data to the cache
const GET_ORGANIZATION_ERROR_REPORTING_CONSENT = gql`
  query GetOrganizationErrorReportingConsent {
    organization {
      displayName
      email
      errorReportingConsent
    }
  }
`;

interface ApolloQueryData {
  organization: Organization;
}

const PromptController: React.FC = () => {
  const { showModal } = useModal();
  const { data } = useQuery<ApolloQueryData>(GET_ORGANIZATION_ERROR_REPORTING_CONSENT);

  React.useEffect(() => {
    if (data?.organization.errorReportingConsent === null) {
      showModal({ modal: MODALS.ANALYTICS_CONSENT });
    }
  }, [data]);
  return null;
};

export default PromptController;
