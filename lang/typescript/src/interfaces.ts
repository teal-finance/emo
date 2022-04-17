import { EmoHook } from './types';

interface EmoParams {
  zone?: string | null;
  activatePrint?: boolean;
  activateEmojis?: boolean;
  hook?: EmoHook | null;
}

export { EmoParams };