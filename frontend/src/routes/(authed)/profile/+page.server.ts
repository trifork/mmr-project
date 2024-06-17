import { redirect } from '@sveltejs/kit';

export const load = async ({ locals: { apiClient } }) => {
  const { userId } = await apiClient.profileApi.v1ProfileGet();

  if (userId != null) {
    redirect(303, `/player/${userId}`);
  }

  return {};
};
