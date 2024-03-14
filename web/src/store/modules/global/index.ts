import { defineStore } from 'pinia';
import piniaStore from '@/store';
import { getRegisterCount, getShareStatistics } from '@/service/api/global';
import { onWarning, onError } from '@/utils/messages';
import useTimer from '@/hooks/useTimer';
import { Notifications } from '@/models/global';
import { useStorage } from '@/utils/use-storage';

export interface GlobalStateProps {
  register_count: number;
  share_count: number;
  click_num: number;
  fetching: boolean;
  notifications: Notifications[];
  has_read_notifications: string[];
}

const initialState = {
  has_read_notifications: useStorage('has_read_notifications'),
};

export const useGlobalStore = defineStore({
  id: 'global',
  state: () =>
    ({
      register_count: -1,
      click_num: -1,
      share_count: -1,
      fetching: false,
      has_read_notifications: initialState.has_read_notifications || [],
    } as GlobalStateProps),
  getters: {
    // register_count: state => state.register_count,
  },
  actions: {
    async onGetRegisterCountAction() {
      try {
        this.fetching = true;
        const { data } = await getRegisterCount();
        if (data.msg === 'success') {
          this.register_count = data.count;
          this.fetching = false;
        } else {
          onWarning(data.msg);
        }
      } catch (error) {
        useTimer(() => {
          this.register_count = 0;
          this.fetching = false;
        }, 3);
        onError(`${error}`);
      }
    },
    onSetNotificationsStatusAction(key: string) {
      if (!this.has_read_notifications.includes(key)) {
        this.has_read_notifications.push(key);
        useStorage('has_read_notifications', this.has_read_notifications);
      }
    },
    async onGetShareStatisticsAction() {
      try {
        this.fetching = true;
        const { data } = await getShareStatistics();
        if (data.msg === 'success') {
          this.click_num = data.click_num;
          this.share_count = data.share_count;
          this.fetching = false;
        } else {
          onWarning(data.msg);
        }
      } catch (error) {
        useTimer(() => {
          this.click_num = 0;
          this.share_count = 0;
          this.fetching = false;
        }, 3);
        onError(`${error}`);
      }
    },
  },
});

export function useGlobalOutsideStore() {
  return useGlobalStore(piniaStore);
}
