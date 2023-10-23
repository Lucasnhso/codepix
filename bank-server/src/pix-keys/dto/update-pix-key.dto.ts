import { PartialType } from '@nestjs/mapped-types';
import { CreatePixKeyDto } from './create-pix-key.dto';
import { PixKeyKind } from '../entities/pix-key.entity';

export class UpdatePixKeyDto extends PartialType(CreatePixKeyDto) {
  key: string;

  kind: PixKeyKind;
}
