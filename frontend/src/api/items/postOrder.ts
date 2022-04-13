import type { UseMutationOptions } from 'react-query';
import { useMutation } from 'react-query';

import { apiClient } from '../apiClient';
import type { IOrderData } from '../types/item';
import type { IRequestException } from '../types/requestException';

interface StatusResponse {
  status: string;
}

export const orderItems = async (data: IOrderData) => {
  const response = await apiClient.post('order', {
    json: { ...data },
  });
  return (await response.json()) as StatusResponse;
};

export const useOrderMutation = (
  options: UseMutationOptions<StatusResponse, IRequestException, IOrderData>
) => {
  return useMutation<StatusResponse, IRequestException, IOrderData>(
    async (data) => await orderItems(data),
    options
  );
};
