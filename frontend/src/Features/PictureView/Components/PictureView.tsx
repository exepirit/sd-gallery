import { Badge, Button, Card, Grid, Image, Text, useTheme } from "@nextui-org/react";
import { ReactElement } from "react";
import { Link as RouterLink } from "react-router-dom";
import { Link } from "@nextui-org/react";
import { Picture } from "../Models";

export type PictureViewProps = {
  picture: Picture
};

export const PictureView = (props: PictureViewProps) => {
  const { theme } = useTheme();

  const renderGenerationInfo = (): ReactElement => {
    const renderProp = (key: string, value: any, isLong?: boolean): ReactElement => (
      <li>
        <Text span>{key}: </Text>
        <Text span={isLong ? !isLong: true}>
          <code>{value}</code>
        </Text>
      </li>
    );

    return (
      <ul>
        {renderProp('Prompt', props.picture.info.prompt, true)}
        {renderProp('Negative prompt', props.picture.info.negativePrompt, true)}
        {renderProp('Steps', props.picture.info.steps)}
        {renderProp('Seed', props.picture.info.seed)}
        {renderProp('Sampler', props.picture.info.sampler)}
        {renderProp('CFG Scale', props.picture.info.cfgScale)}
      </ul>
    );
  };

  return (
    <Grid.Container>
      <Grid md={9} css={{
        dflex: 'center',
        ai: 'center'
      }}>
        <Image src={`/api/images/${props.picture.id}`}
          width={props.picture.size.width}
          height={props.picture.size.height}
          autoResize/>
      </Grid>
      <Grid md={3} direction='column' css={{
        h: '100vh',
        p: 36,
      }}>
        <div>
          {props.picture.tags.map((tag, idx) => <Badge key={idx}>{tag}</Badge>)}
        </div>
        {renderGenerationInfo()}
        <Card variant='flat'>
          <Card.Body css={{
            display: 'flex',
            flexDirection: 'row',
            justifyContent: 'space-between'
          }}>
            <div>
              <Text>Stable Diffusion model</Text>
              <Text color={theme?.colors.accents7.value}>{props.picture.info.modelHash}</Text>
            </div>
            <div>
              <Badge color='primary'>Model</Badge>
            </div>
          </Card.Body>
        </Card>
      </Grid>
    </Grid.Container>
  )
};