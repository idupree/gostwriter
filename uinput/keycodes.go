/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package uinput

/*
  #include <linux/input.h>
*/
import "C"

const (
    KEY_RESERVED          = C.KEY_RESERVED             /* 0     */
    KEY_ESC               = C.KEY_ESC                  /* 1     */
    KEY_1                 = C.KEY_1                    /* 2     */
    KEY_2                 = C.KEY_2                    /* 3     */
    KEY_3                 = C.KEY_3                    /* 4     */
    KEY_4                 = C.KEY_4                    /* 5     */
    KEY_5                 = C.KEY_5                    /* 6     */
    KEY_6                 = C.KEY_6                    /* 7     */
    KEY_7                 = C.KEY_7                    /* 8     */
    KEY_8                 = C.KEY_8                    /* 9     */
    KEY_9                 = C.KEY_9                    /* 10    */
    KEY_0                 = C.KEY_0                    /* 11    */
    KEY_MINUS             = C.KEY_MINUS                /* 12    */
    KEY_EQUAL             = C.KEY_EQUAL                /* 13    */
    KEY_BACKSPACE         = C.KEY_BACKSPACE            /* 14    */
    KEY_TAB               = C.KEY_TAB                  /* 15    */
    KEY_Q                 = C.KEY_Q                    /* 16    */
    KEY_W                 = C.KEY_W                    /* 17    */
    KEY_E                 = C.KEY_E                    /* 18    */
    KEY_R                 = C.KEY_R                    /* 19    */
    KEY_T                 = C.KEY_T                    /* 20    */
    KEY_Y                 = C.KEY_Y                    /* 21    */
    KEY_U                 = C.KEY_U                    /* 22    */
    KEY_I                 = C.KEY_I                    /* 23    */
    KEY_O                 = C.KEY_O                    /* 24    */
    KEY_P                 = C.KEY_P                    /* 25    */
    KEY_LEFTBRACE         = C.KEY_LEFTBRACE            /* 26    */
    KEY_RIGHTBRACE        = C.KEY_RIGHTBRACE           /* 27    */
    KEY_ENTER             = C.KEY_ENTER                /* 28    */
    KEY_LEFTCTRL          = C.KEY_LEFTCTRL             /* 29    */
    KEY_A                 = C.KEY_A                    /* 30    */
    KEY_S                 = C.KEY_S                    /* 31    */
    KEY_D                 = C.KEY_D                    /* 32    */
    KEY_F                 = C.KEY_F                    /* 33    */
    KEY_G                 = C.KEY_G                    /* 34    */
    KEY_H                 = C.KEY_H                    /* 35    */
    KEY_J                 = C.KEY_J                    /* 36    */
    KEY_K                 = C.KEY_K                    /* 37    */
    KEY_L                 = C.KEY_L                    /* 38    */
    KEY_SEMICOLON         = C.KEY_SEMICOLON            /* 39    */
    KEY_APOSTROPHE        = C.KEY_APOSTROPHE           /* 40    */
    KEY_GRAVE             = C.KEY_GRAVE                /* 41    */
    KEY_LEFTSHIFT         = C.KEY_LEFTSHIFT            /* 42    */
    KEY_BACKSLASH         = C.KEY_BACKSLASH            /* 43    */
    KEY_Z                 = C.KEY_Z                    /* 44    */
    KEY_X                 = C.KEY_X                    /* 45    */
    KEY_C                 = C.KEY_C                    /* 46    */
    KEY_V                 = C.KEY_V                    /* 47    */
    KEY_B                 = C.KEY_B                    /* 48    */
    KEY_N                 = C.KEY_N                    /* 49    */
    KEY_M                 = C.KEY_M                    /* 50    */
    KEY_COMMA             = C.KEY_COMMA                /* 51    */
    KEY_DOT               = C.KEY_DOT                  /* 52    */
    KEY_SLASH             = C.KEY_SLASH                /* 53    */
    KEY_RIGHTSHIFT        = C.KEY_RIGHTSHIFT           /* 54    */
    KEY_KPASTERISK        = C.KEY_KPASTERISK           /* 55    */
    KEY_LEFTALT           = C.KEY_LEFTALT              /* 56    */
    KEY_SPACE             = C.KEY_SPACE                /* 57    */
    KEY_CAPSLOCK          = C.KEY_CAPSLOCK             /* 58    */
    KEY_F1                = C.KEY_F1                   /* 59    */
    KEY_F2                = C.KEY_F2                   /* 60    */
    KEY_F3                = C.KEY_F3                   /* 61    */
    KEY_F4                = C.KEY_F4                   /* 62    */
    KEY_F5                = C.KEY_F5                   /* 63    */
    KEY_F6                = C.KEY_F6                   /* 64    */
    KEY_F7                = C.KEY_F7                   /* 65    */
    KEY_F8                = C.KEY_F8                   /* 66    */
    KEY_F9                = C.KEY_F9                   /* 67    */
    KEY_F10               = C.KEY_F10                  /* 68    */
    KEY_NUMLOCK           = C.KEY_NUMLOCK              /* 69    */
    KEY_SCROLLLOCK        = C.KEY_SCROLLLOCK           /* 70    */
    KEY_KP7               = C.KEY_KP7                  /* 71    */
    KEY_KP8               = C.KEY_KP8                  /* 72    */
    KEY_KP9               = C.KEY_KP9                  /* 73    */
    KEY_KPMINUS           = C.KEY_KPMINUS              /* 74    */
    KEY_KP4               = C.KEY_KP4                  /* 75    */
    KEY_KP5               = C.KEY_KP5                  /* 76    */
    KEY_KP6               = C.KEY_KP6                  /* 77    */
    KEY_KPPLUS            = C.KEY_KPPLUS               /* 78    */
    KEY_KP1               = C.KEY_KP1                  /* 79    */
    KEY_KP2               = C.KEY_KP2                  /* 80    */
    KEY_KP3               = C.KEY_KP3                  /* 81    */
    KEY_KP0               = C.KEY_KP0                  /* 82    */
    KEY_KPDOT             = C.KEY_KPDOT                /* 83    */
    KEY_ZENKAKUHANKAKU    = C.KEY_ZENKAKUHANKAKU       /* 85    */
    KEY_102ND             = C.KEY_102ND                /* 86    */
    KEY_F11               = C.KEY_F11                  /* 87    */
    KEY_F12               = C.KEY_F12                  /* 88    */
    KEY_RO                = C.KEY_RO                   /* 89    */
    KEY_KATAKANA          = C.KEY_KATAKANA             /* 90    */
    KEY_HIRAGANA          = C.KEY_HIRAGANA             /* 91    */
    KEY_HENKAN            = C.KEY_HENKAN               /* 92    */
    KEY_KATAKANAHIRAGANA  = C.KEY_KATAKANAHIRAGANA     /* 93    */
    KEY_MUHENKAN          = C.KEY_MUHENKAN             /* 94    */
    KEY_KPJPCOMMA         = C.KEY_KPJPCOMMA            /* 95    */
    KEY_KPENTER           = C.KEY_KPENTER              /* 96    */
    KEY_RIGHTCTRL         = C.KEY_RIGHTCTRL            /* 97    */
    KEY_KPSLASH           = C.KEY_KPSLASH              /* 98    */
    KEY_SYSRQ             = C.KEY_SYSRQ                /* 99    */
    KEY_RIGHTALT          = C.KEY_RIGHTALT             /* 100   */
    KEY_LINEFEED          = C.KEY_LINEFEED             /* 101   */
    KEY_HOME              = C.KEY_HOME                 /* 102   */
    KEY_UP                = C.KEY_UP                   /* 103   */
    KEY_PAGEUP            = C.KEY_PAGEUP               /* 104   */
    KEY_LEFT              = C.KEY_LEFT                 /* 105   */
    KEY_RIGHT             = C.KEY_RIGHT                /* 106   */
    KEY_END               = C.KEY_END                  /* 107   */
    KEY_DOWN              = C.KEY_DOWN                 /* 108   */
    KEY_PAGEDOWN          = C.KEY_PAGEDOWN             /* 109   */
    KEY_INSERT            = C.KEY_INSERT               /* 110   */
    KEY_DELETE            = C.KEY_DELETE               /* 111   */
    KEY_MACRO             = C.KEY_MACRO                /* 112   */
    KEY_MUTE              = C.KEY_MUTE                 /* 113   */
    KEY_VOLUMEDOWN        = C.KEY_VOLUMEDOWN           /* 114   */
    KEY_VOLUMEUP          = C.KEY_VOLUMEUP             /* 115   */
    KEY_POWER             = C.KEY_POWER                /* 116   */
    KEY_KPEQUAL           = C.KEY_KPEQUAL              /* 117   */
    KEY_KPPLUSMINUS       = C.KEY_KPPLUSMINUS          /* 118   */
    KEY_PAUSE             = C.KEY_PAUSE                /* 119   */
    KEY_SCALE             = C.KEY_SCALE                /* 120   */
    KEY_KPCOMMA           = C.KEY_KPCOMMA              /* 121   */
    KEY_HANGEUL           = C.KEY_HANGEUL              /* 122   */
    KEY_HANGUEL           = C.KEY_HANGUEL              /* KEY_HANGEUL */
    KEY_HANJA             = C.KEY_HANJA                /* 123   */
    KEY_YEN               = C.KEY_YEN                  /* 124   */
    KEY_LEFTMETA          = C.KEY_LEFTMETA             /* 125   */
    KEY_RIGHTMETA         = C.KEY_RIGHTMETA            /* 126   */
    KEY_COMPOSE           = C.KEY_COMPOSE              /* 127   */
    KEY_STOP              = C.KEY_STOP                 /* 128   */
    KEY_AGAIN             = C.KEY_AGAIN                /* 129   */
    KEY_PROPS             = C.KEY_PROPS                /* 130   */
    KEY_UNDO              = C.KEY_UNDO                 /* 131   */
    KEY_FRONT             = C.KEY_FRONT                /* 132   */
    KEY_COPY              = C.KEY_COPY                 /* 133   */
    KEY_OPEN              = C.KEY_OPEN                 /* 134   */
    KEY_PASTE             = C.KEY_PASTE                /* 135   */
    KEY_FIND              = C.KEY_FIND                 /* 136   */
    KEY_CUT               = C.KEY_CUT                  /* 137   */
    KEY_HELP              = C.KEY_HELP                 /* 138   */
    KEY_MENU              = C.KEY_MENU                 /* 139   */
    KEY_CALC              = C.KEY_CALC                 /* 140   */
    KEY_SETUP             = C.KEY_SETUP                /* 141   */
    KEY_SLEEP             = C.KEY_SLEEP                /* 142   */
    KEY_WAKEUP            = C.KEY_WAKEUP               /* 143   */
    KEY_FILE              = C.KEY_FILE                 /* 144   */
    KEY_SENDFILE          = C.KEY_SENDFILE             /* 145   */
    KEY_DELETEFILE        = C.KEY_DELETEFILE           /* 146   */
    KEY_XFER              = C.KEY_XFER                 /* 147   */
    KEY_PROG1             = C.KEY_PROG1                /* 148   */
    KEY_PROG2             = C.KEY_PROG2                /* 149   */
    KEY_WWW               = C.KEY_WWW                  /* 150   */
    KEY_MSDOS             = C.KEY_MSDOS                /* 151   */
    KEY_COFFEE            = C.KEY_COFFEE               /* 152   */
    KEY_SCREENLOCK        = C.KEY_SCREENLOCK           /* KEY_COFFEE */
    KEY_DIRECTION         = C.KEY_DIRECTION            /* 153   */
    KEY_CYCLEWINDOWS      = C.KEY_CYCLEWINDOWS         /* 154   */
    KEY_MAIL              = C.KEY_MAIL                 /* 155   */
    KEY_BOOKMARKS         = C.KEY_BOOKMARKS            /* 156   */
    KEY_COMPUTER          = C.KEY_COMPUTER             /* 157   */
    KEY_BACK              = C.KEY_BACK                 /* 158   */
    KEY_FORWARD           = C.KEY_FORWARD              /* 159   */
    KEY_CLOSECD           = C.KEY_CLOSECD              /* 160   */
    KEY_EJECTCD           = C.KEY_EJECTCD              /* 161   */
    KEY_EJECTCLOSECD      = C.KEY_EJECTCLOSECD         /* 162   */
    KEY_NEXTSONG          = C.KEY_NEXTSONG             /* 163   */
    KEY_PLAYPAUSE         = C.KEY_PLAYPAUSE            /* 164   */
    KEY_PREVIOUSSONG      = C.KEY_PREVIOUSSONG         /* 165   */
    KEY_STOPCD            = C.KEY_STOPCD               /* 166   */
    KEY_RECORD            = C.KEY_RECORD               /* 167   */
    KEY_REWIND            = C.KEY_REWIND               /* 168   */
    KEY_PHONE             = C.KEY_PHONE                /* 169   */
    KEY_ISO               = C.KEY_ISO                  /* 170   */
    KEY_CONFIG            = C.KEY_CONFIG               /* 171   */
    KEY_HOMEPAGE          = C.KEY_HOMEPAGE             /* 172   */
    KEY_REFRESH           = C.KEY_REFRESH              /* 173   */
    KEY_EXIT              = C.KEY_EXIT                 /* 174   */
    KEY_MOVE              = C.KEY_MOVE                 /* 175   */
    KEY_EDIT              = C.KEY_EDIT                 /* 176   */
    KEY_SCROLLUP          = C.KEY_SCROLLUP             /* 177   */
    KEY_SCROLLDOWN        = C.KEY_SCROLLDOWN           /* 178   */
    KEY_KPLEFTPAREN       = C.KEY_KPLEFTPAREN          /* 179   */
    KEY_KPRIGHTPAREN      = C.KEY_KPRIGHTPAREN         /* 180   */
    KEY_NEW               = C.KEY_NEW                  /* 181   */
    KEY_REDO              = C.KEY_REDO                 /* 182   */
    KEY_F13               = C.KEY_F13                  /* 183   */
    KEY_F14               = C.KEY_F14                  /* 184   */
    KEY_F15               = C.KEY_F15                  /* 185   */
    KEY_F16               = C.KEY_F16                  /* 186   */
    KEY_F17               = C.KEY_F17                  /* 187   */
    KEY_F18               = C.KEY_F18                  /* 188   */
    KEY_F19               = C.KEY_F19                  /* 189   */
    KEY_F20               = C.KEY_F20                  /* 190   */
    KEY_F21               = C.KEY_F21                  /* 191   */
    KEY_F22               = C.KEY_F22                  /* 192   */
    KEY_F23               = C.KEY_F23                  /* 193   */
    KEY_F24               = C.KEY_F24                  /* 194   */
    KEY_PLAYCD            = C.KEY_PLAYCD               /* 200   */
    KEY_PAUSECD           = C.KEY_PAUSECD              /* 201   */
    KEY_PROG3             = C.KEY_PROG3                /* 202   */
    KEY_PROG4             = C.KEY_PROG4                /* 203   */
    KEY_DASHBOARD         = C.KEY_DASHBOARD            /* 204   */
    KEY_SUSPEND           = C.KEY_SUSPEND              /* 205   */
    KEY_CLOSE             = C.KEY_CLOSE                /* 206   */
    KEY_PLAY              = C.KEY_PLAY                 /* 207   */
    KEY_FASTFORWARD       = C.KEY_FASTFORWARD          /* 208   */
    KEY_BASSBOOST         = C.KEY_BASSBOOST            /* 209   */
    KEY_PRINT             = C.KEY_PRINT                /* 210   */
    KEY_HP                = C.KEY_HP                   /* 211   */
    KEY_CAMERA            = C.KEY_CAMERA               /* 212   */
    KEY_SOUND             = C.KEY_SOUND                /* 213   */
    KEY_QUESTION          = C.KEY_QUESTION             /* 214   */
    KEY_EMAIL             = C.KEY_EMAIL                /* 215   */
    KEY_CHAT              = C.KEY_CHAT                 /* 216   */
    KEY_SEARCH            = C.KEY_SEARCH               /* 217   */
    KEY_CONNECT           = C.KEY_CONNECT              /* 218   */
    KEY_FINANCE           = C.KEY_FINANCE              /* 219   */
    KEY_SPORT             = C.KEY_SPORT                /* 220   */
    KEY_SHOP              = C.KEY_SHOP                 /* 221   */
    KEY_ALTERASE          = C.KEY_ALTERASE             /* 222   */
    KEY_CANCEL            = C.KEY_CANCEL               /* 223   */
    KEY_BRIGHTNESSDOWN    = C.KEY_BRIGHTNESSDOWN       /* 224   */
    KEY_BRIGHTNESSUP      = C.KEY_BRIGHTNESSUP         /* 225   */
    KEY_MEDIA             = C.KEY_MEDIA                /* 226   */
    KEY_SWITCHVIDEOMODE   = C.KEY_SWITCHVIDEOMODE      /* 227   */
    KEY_KBDILLUMTOGGLE    = C.KEY_KBDILLUMTOGGLE       /* 228   */
    KEY_KBDILLUMDOWN      = C.KEY_KBDILLUMDOWN         /* 229   */
    KEY_KBDILLUMUP        = C.KEY_KBDILLUMUP           /* 230   */
    KEY_SEND              = C.KEY_SEND                 /* 231   */
    KEY_REPLY             = C.KEY_REPLY                /* 232   */
    KEY_FORWARDMAIL       = C.KEY_FORWARDMAIL          /* 233   */
    KEY_SAVE              = C.KEY_SAVE                 /* 234   */
    KEY_DOCUMENTS         = C.KEY_DOCUMENTS            /* 235   */
    KEY_BATTERY           = C.KEY_BATTERY              /* 236   */
    KEY_BLUETOOTH         = C.KEY_BLUETOOTH            /* 237   */
    KEY_WLAN              = C.KEY_WLAN                 /* 238   */
    KEY_UWB               = C.KEY_UWB                  /* 239   */
    KEY_UNKNOWN           = C.KEY_UNKNOWN              /* 240   */
    KEY_VIDEO_NEXT        = C.KEY_VIDEO_NEXT           /* 241   */
    KEY_VIDEO_PREV        = C.KEY_VIDEO_PREV           /* 242   */
    KEY_BRIGHTNESS_CYCLE  = C.KEY_BRIGHTNESS_CYCLE     /* 243   */
    KEY_BRIGHTNESS_AUTO   = C.KEY_BRIGHTNESS_AUTO      /* 244   */
    KEY_BRIGHTNESS_ZERO   = C.KEY_BRIGHTNESS_ZERO      /* KEY_BRIGHTNESS_AUTO */
    KEY_DISPLAY_OFF       = C.KEY_DISPLAY_OFF          /* 245   */
    KEY_WWAN              = C.KEY_WWAN                 /* 246   */
    KEY_WIMAX             = C.KEY_WIMAX                /* KEY_WWAN */
    KEY_RFKILL            = C.KEY_RFKILL               /* 247   */
    KEY_MICMUTE           = C.KEY_MICMUTE              /* 248   */
    KEY_OK                = C.KEY_OK                   /* 0x160 */
    KEY_SELECT            = C.KEY_SELECT               /* 0x161 */
    KEY_GOTO              = C.KEY_GOTO                 /* 0x162 */
    KEY_CLEAR             = C.KEY_CLEAR                /* 0x163 */
    KEY_POWER2            = C.KEY_POWER2               /* 0x164 */
    KEY_OPTION            = C.KEY_OPTION               /* 0x165 */
    KEY_INFO              = C.KEY_INFO                 /* 0x166 */
    KEY_TIME              = C.KEY_TIME                 /* 0x167 */
    KEY_VENDOR            = C.KEY_VENDOR               /* 0x168 */
    KEY_ARCHIVE           = C.KEY_ARCHIVE              /* 0x169 */
    KEY_PROGRAM           = C.KEY_PROGRAM              /* 0x16a */
    KEY_CHANNEL           = C.KEY_CHANNEL              /* 0x16b */
    KEY_FAVORITES         = C.KEY_FAVORITES            /* 0x16c */
    KEY_EPG               = C.KEY_EPG                  /* 0x16d */
    KEY_PVR               = C.KEY_PVR                  /* 0x16e */
    KEY_MHP               = C.KEY_MHP                  /* 0x16f */
    KEY_LANGUAGE          = C.KEY_LANGUAGE             /* 0x170 */
    KEY_TITLE             = C.KEY_TITLE                /* 0x171 */
    KEY_SUBTITLE          = C.KEY_SUBTITLE             /* 0x172 */
    KEY_ANGLE             = C.KEY_ANGLE                /* 0x173 */
    KEY_ZOOM              = C.KEY_ZOOM                 /* 0x174 */
    KEY_MODE              = C.KEY_MODE                 /* 0x175 */
    KEY_KEYBOARD          = C.KEY_KEYBOARD             /* 0x176 */
    KEY_SCREEN            = C.KEY_SCREEN               /* 0x177 */
    KEY_PC                = C.KEY_PC                   /* 0x178 */
    KEY_TV                = C.KEY_TV                   /* 0x179 */
    KEY_TV2               = C.KEY_TV2                  /* 0x17a */
    KEY_VCR               = C.KEY_VCR                  /* 0x17b */
    KEY_VCR2              = C.KEY_VCR2                 /* 0x17c */
    KEY_SAT               = C.KEY_SAT                  /* 0x17d */
    KEY_SAT2              = C.KEY_SAT2                 /* 0x17e */
    KEY_CD                = C.KEY_CD                   /* 0x17f */
    KEY_TAPE              = C.KEY_TAPE                 /* 0x180 */
    KEY_RADIO             = C.KEY_RADIO                /* 0x181 */
    KEY_TUNER             = C.KEY_TUNER                /* 0x182 */
    KEY_PLAYER            = C.KEY_PLAYER               /* 0x183 */
    KEY_TEXT              = C.KEY_TEXT                 /* 0x184 */
    KEY_DVD               = C.KEY_DVD                  /* 0x185 */
    KEY_AUX               = C.KEY_AUX                  /* 0x186 */
    KEY_MP3               = C.KEY_MP3                  /* 0x187 */
    KEY_AUDIO             = C.KEY_AUDIO                /* 0x188 */
    KEY_VIDEO             = C.KEY_VIDEO                /* 0x189 */
    KEY_DIRECTORY         = C.KEY_DIRECTORY            /* 0x18a */
    KEY_LIST              = C.KEY_LIST                 /* 0x18b */
    KEY_MEMO              = C.KEY_MEMO                 /* 0x18c */
    KEY_CALENDAR          = C.KEY_CALENDAR             /* 0x18d */
    KEY_RED               = C.KEY_RED                  /* 0x18e */
    KEY_GREEN             = C.KEY_GREEN                /* 0x18f */
    KEY_YELLOW            = C.KEY_YELLOW               /* 0x190 */
    KEY_BLUE              = C.KEY_BLUE                 /* 0x191 */
    KEY_CHANNELUP         = C.KEY_CHANNELUP            /* 0x192 */
    KEY_CHANNELDOWN       = C.KEY_CHANNELDOWN          /* 0x193 */
    KEY_FIRST             = C.KEY_FIRST                /* 0x194 */
    KEY_LAST              = C.KEY_LAST                 /* 0x195 */
    KEY_AB                = C.KEY_AB                   /* 0x196 */
    KEY_NEXT              = C.KEY_NEXT                 /* 0x197 */
    KEY_RESTART           = C.KEY_RESTART              /* 0x198 */
    KEY_SLOW              = C.KEY_SLOW                 /* 0x199 */
    KEY_SHUFFLE           = C.KEY_SHUFFLE              /* 0x19a */
    KEY_BREAK             = C.KEY_BREAK                /* 0x19b */
    KEY_PREVIOUS          = C.KEY_PREVIOUS             /* 0x19c */
    KEY_DIGITS            = C.KEY_DIGITS               /* 0x19d */
    KEY_TEEN              = C.KEY_TEEN                 /* 0x19e */
    KEY_TWEN              = C.KEY_TWEN                 /* 0x19f */
    KEY_VIDEOPHONE        = C.KEY_VIDEOPHONE           /* 0x1a0 */
    KEY_GAMES             = C.KEY_GAMES                /* 0x1a1 */
    KEY_ZOOMIN            = C.KEY_ZOOMIN               /* 0x1a2 */
    KEY_ZOOMOUT           = C.KEY_ZOOMOUT              /* 0x1a3 */
    KEY_ZOOMRESET         = C.KEY_ZOOMRESET            /* 0x1a4 */
    KEY_WORDPROCESSOR     = C.KEY_WORDPROCESSOR        /* 0x1a5 */
    KEY_EDITOR            = C.KEY_EDITOR               /* 0x1a6 */
    KEY_SPREADSHEET       = C.KEY_SPREADSHEET          /* 0x1a7 */
    KEY_GRAPHICSEDITOR    = C.KEY_GRAPHICSEDITOR       /* 0x1a8 */
    KEY_PRESENTATION      = C.KEY_PRESENTATION         /* 0x1a9 */
    KEY_DATABASE          = C.KEY_DATABASE             /* 0x1aa */
    KEY_NEWS              = C.KEY_NEWS                 /* 0x1ab */
    KEY_VOICEMAIL         = C.KEY_VOICEMAIL            /* 0x1ac */
    KEY_ADDRESSBOOK       = C.KEY_ADDRESSBOOK          /* 0x1ad */
    KEY_MESSENGER         = C.KEY_MESSENGER            /* 0x1ae */
    KEY_DISPLAYTOGGLE     = C.KEY_DISPLAYTOGGLE        /* 0x1af */
    KEY_BRIGHTNESS_TOGGLE = C.KEY_BRIGHTNESS_TOGGLE    /* KEY_DISPLAYTOGGLE */
    KEY_SPELLCHECK        = C.KEY_SPELLCHECK           /* 0x1b0 */
    KEY_LOGOFF            = C.KEY_LOGOFF               /* 0x1b1 */
    KEY_DOLLAR            = C.KEY_DOLLAR               /* 0x1b2 */
    KEY_EURO              = C.KEY_EURO                 /* 0x1b3 */
    KEY_FRAMEBACK         = C.KEY_FRAMEBACK            /* 0x1b4 */
    KEY_FRAMEFORWARD      = C.KEY_FRAMEFORWARD         /* 0x1b5 */
    KEY_CONTEXT_MENU      = C.KEY_CONTEXT_MENU         /* 0x1b6 */
    KEY_MEDIA_REPEAT      = C.KEY_MEDIA_REPEAT         /* 0x1b7 */
    KEY_10CHANNELSUP      = C.KEY_10CHANNELSUP         /* 0x1b8 */
    KEY_10CHANNELSDOWN    = C.KEY_10CHANNELSDOWN       /* 0x1b9 */
    KEY_IMAGES            = C.KEY_IMAGES               /* 0x1ba */
    KEY_DEL_EOL           = C.KEY_DEL_EOL              /* 0x1c0 */
    KEY_DEL_EOS           = C.KEY_DEL_EOS              /* 0x1c1 */
    KEY_INS_LINE          = C.KEY_INS_LINE             /* 0x1c2 */
    KEY_DEL_LINE          = C.KEY_DEL_LINE             /* 0x1c3 */
    KEY_FN                = C.KEY_FN                   /* 0x1d0 */
    KEY_FN_ESC            = C.KEY_FN_ESC               /* 0x1d1 */
    KEY_FN_F1             = C.KEY_FN_F1                /* 0x1d2 */
    KEY_FN_F2             = C.KEY_FN_F2                /* 0x1d3 */
    KEY_FN_F3             = C.KEY_FN_F3                /* 0x1d4 */
    KEY_FN_F4             = C.KEY_FN_F4                /* 0x1d5 */
    KEY_FN_F5             = C.KEY_FN_F5                /* 0x1d6 */
    KEY_FN_F6             = C.KEY_FN_F6                /* 0x1d7 */
    KEY_FN_F7             = C.KEY_FN_F7                /* 0x1d8 */
    KEY_FN_F8             = C.KEY_FN_F8                /* 0x1d9 */
    KEY_FN_F9             = C.KEY_FN_F9                /* 0x1da */
    KEY_FN_F10            = C.KEY_FN_F10               /* 0x1db */
    KEY_FN_F11            = C.KEY_FN_F11               /* 0x1dc */
    KEY_FN_F12            = C.KEY_FN_F12               /* 0x1dd */
    KEY_FN_1              = C.KEY_FN_1                 /* 0x1de */
    KEY_FN_2              = C.KEY_FN_2                 /* 0x1df */
    KEY_FN_D              = C.KEY_FN_D                 /* 0x1e0 */
    KEY_FN_E              = C.KEY_FN_E                 /* 0x1e1 */
    KEY_FN_F              = C.KEY_FN_F                 /* 0x1e2 */
    KEY_FN_S              = C.KEY_FN_S                 /* 0x1e3 */
    KEY_FN_B              = C.KEY_FN_B                 /* 0x1e4 */
    KEY_BRL_DOT1          = C.KEY_BRL_DOT1             /* 0x1f1 */
    KEY_BRL_DOT2          = C.KEY_BRL_DOT2             /* 0x1f2 */
    KEY_BRL_DOT3          = C.KEY_BRL_DOT3             /* 0x1f3 */
    KEY_BRL_DOT4          = C.KEY_BRL_DOT4             /* 0x1f4 */
    KEY_BRL_DOT5          = C.KEY_BRL_DOT5             /* 0x1f5 */
    KEY_BRL_DOT6          = C.KEY_BRL_DOT6             /* 0x1f6 */
    KEY_BRL_DOT7          = C.KEY_BRL_DOT7             /* 0x1f7 */
    KEY_BRL_DOT8          = C.KEY_BRL_DOT8             /* 0x1f8 */
    KEY_BRL_DOT9          = C.KEY_BRL_DOT9             /* 0x1f9 */
    KEY_BRL_DOT10         = C.KEY_BRL_DOT10            /* 0x1fa */
    KEY_NUMERIC_0         = C.KEY_NUMERIC_0            /* 0x200 */
    KEY_NUMERIC_1         = C.KEY_NUMERIC_1            /* 0x201 */
    KEY_NUMERIC_2         = C.KEY_NUMERIC_2            /* 0x202 */
    KEY_NUMERIC_3         = C.KEY_NUMERIC_3            /* 0x203 */
    KEY_NUMERIC_4         = C.KEY_NUMERIC_4            /* 0x204 */
    KEY_NUMERIC_5         = C.KEY_NUMERIC_5            /* 0x205 */
    KEY_NUMERIC_6         = C.KEY_NUMERIC_6            /* 0x206 */
    KEY_NUMERIC_7         = C.KEY_NUMERIC_7            /* 0x207 */
    KEY_NUMERIC_8         = C.KEY_NUMERIC_8            /* 0x208 */
    KEY_NUMERIC_9         = C.KEY_NUMERIC_9            /* 0x209 */
    KEY_NUMERIC_STAR      = C.KEY_NUMERIC_STAR         /* 0x20a */
    KEY_NUMERIC_POUND     = C.KEY_NUMERIC_POUND        /* 0x20b */
    KEY_CAMERA_FOCUS      = C.KEY_CAMERA_FOCUS         /* 0x210 */
    KEY_WPS_BUTTON        = C.KEY_WPS_BUTTON           /* 0x211 */
    KEY_TOUCHPAD_TOGGLE   = C.KEY_TOUCHPAD_TOGGLE      /* 0x212 */
    KEY_TOUCHPAD_ON       = C.KEY_TOUCHPAD_ON          /* 0x213 */
    KEY_TOUCHPAD_OFF      = C.KEY_TOUCHPAD_OFF         /* 0x214 */
    KEY_CAMERA_ZOOMIN     = C.KEY_CAMERA_ZOOMIN        /* 0x215 */
    KEY_CAMERA_ZOOMOUT    = C.KEY_CAMERA_ZOOMOUT       /* 0x216 */
    KEY_CAMERA_UP         = C.KEY_CAMERA_UP            /* 0x217 */
    KEY_CAMERA_DOWN       = C.KEY_CAMERA_DOWN          /* 0x218 */
    KEY_CAMERA_LEFT       = C.KEY_CAMERA_LEFT          /* 0x219 */
    KEY_CAMERA_RIGHT      = C.KEY_CAMERA_RIGHT         /* 0x21a */
    KEY_ATTENDANT_ON      = C.KEY_ATTENDANT_ON         /* 0x21b */
    KEY_ATTENDANT_OFF     = C.KEY_ATTENDANT_OFF        /* 0x21c */
    KEY_ATTENDANT_TOGGLE  = C.KEY_ATTENDANT_TOGGLE     /* 0x21d */
    KEY_LIGHTS_TOGGLE     = C.KEY_LIGHTS_TOGGLE        /* 0x21e */
    KEY_ALS_TOGGLE        = C.KEY_ALS_TOGGLE           /* 0x230 */
    KEY_BUTTONCONFIG      = C.KEY_BUTTONCONFIG         /* 0x240 */
    KEY_TASKMANAGER       = C.KEY_TASKMANAGER          /* 0x241 */
    KEY_JOURNAL           = C.KEY_JOURNAL              /* 0x242 */
    KEY_CONTROLPANEL      = C.KEY_CONTROLPANEL         /* 0x243 */
    KEY_APPSELECT         = C.KEY_APPSELECT            /* 0x244 */
    KEY_SCREENSAVER       = C.KEY_SCREENSAVER          /* 0x245 */
    KEY_VOICECOMMAND      = C.KEY_VOICECOMMAND         /* 0x246 */
    KEY_BRIGHTNESS_MIN    = C.KEY_BRIGHTNESS_MIN       /* 0x250 */
    KEY_BRIGHTNESS_MAX    = C.KEY_BRIGHTNESS_MAX       /* 0x251 */
    KEY_MIN_INTERESTING   = C.KEY_MIN_INTERESTING      /* KEY_MUTE */
    KEY_MAX               = C.KEY_MAX                  /* 0x2ff */
    KEY_CNT               = C.KEY_CNT                  /* (KEY_MAX+1) */
)
